package rbTree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

var (
	ErrInexistItem = fmt.Errorf("delete op error: no such item")
)

type Color bool

const (
	BLACK, RED Color = false, true
)

type Node[T constraints.Ordered] struct {
	Key   T
	Color Color
	Left  *Node[T]
	Right *Node[T]
}

func isRed[T constraints.Ordered](h *Node[T]) bool {
	if h == nil {
		return false
	}
	return h.Color == RED
}

func isBlack[T constraints.Ordered](h *Node[T]) bool {
	return !isRed(h)
}

type rbTree[T constraints.Ordered] struct {
	root *Node[T]
}

func NewTree[T constraints.Ordered]() *rbTree[T] {
	return &rbTree[T]{}
}

func (tree *rbTree[T]) Has(key T) bool {
	node := tree.root
	for node != nil {
		if key == node.Key {
			return true
		} else if key < node.Key {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return false
}

func rotateLeft[T constraints.Ordered](h *Node[T]) *Node[T] {
	r := h.Right
	h.Right = r.Left
	r.Left = h
	r.Color = r.Left.Color
	r.Left.Color = RED
	return r
}

func rotateRight[T constraints.Ordered](h *Node[T]) *Node[T] {
	l := h.Left
	h.Left = l.Right
	l.Right = h
	l.Color = l.Right.Color
	l.Right.Color = RED
	return l
}

// colorFlip guarded by the algo design. No need to do null check.
func (n *Node[T]) colorFlip() {
	n.Color = !n.Color
	n.Left.Color = !n.Left.Color
	n.Right.Color = !n.Right.Color
}

func (tree *rbTree[T]) InsertOrReplace(key T) {
	tree.root = insertOrReplace(tree.root, key)
	tree.root.Color = BLACK
}

func insertOrReplace[T constraints.Ordered](h *Node[T], key T) *Node[T] {
	if h == nil { // insert at bottom
		return &Node[T]{
			Key:   key,
			Color: RED,
		}
	}
	if key == h.Key {
		// TODO: replace

	} else if key < h.Key {
		h.Left = insertOrReplace(h.Left, key)
	} else if key > h.Key {
		h.Right = insertOrReplace(h.Right, key)
	}

	// path of way down completed. Now go up.
	h = h.fixUp()
	return h
}

func (tree *rbTree[T]) Delete(key T) error {
	var success bool
	tree.root, success = delete(tree.root, key)
	if tree.root != nil {
		tree.root.Color = BLACK
	}
	if !success {
		return ErrInexistItem
	} else {
		return nil
	}
}

func getMin[T constraints.Ordered](h *Node[T]) *Node[T] {
	if h.Left != nil {
		return getMin(h.Left)
	}
	return h
}

func deleteMin[T constraints.Ordered](h *Node[T]) *Node[T] {
	if h.Left == nil {
		return nil
	}
	if isBlack(h.Left) && isBlack(h.Left.Left) {
		h = moveRedLeft(h)
	}

	h.Left = deleteMin(h.Left)

	return h.fixUp()
}

func delete[T constraints.Ordered](h *Node[T], key T) (*Node[T], bool) {
	var deleted bool
	// if h == nil {
	// 	return nil, true
	// }
	if key < h.Key {
		// if h.left == nil { // needed ?
		// 	return h, false
		// }
		if isBlack(h.Left) && isBlack(h.Left.Left) {
			h = moveRedLeft(h)
		}
		h.Left, deleted = delete(h.Left, key)
	} else {
		if isRed(h.Left) {
			h = rotateRight(h)
		}
		if key == h.Key && h.Right == nil {
			return nil, true
		}
		if isBlack(h.Right) && isBlack(h.Right.Left) {
			h = moveRedRight(h)
		}
		if key == h.Key {
			minNodeInRightSubTree := getMin(h.Right)
			h.Key = minNodeInRightSubTree.Key
			h.Right = deleteMin(h.Right)
		} else {
			h.Right, deleted = delete(h.Right, key)
		}
	}
	return h.fixUp(), deleted
}

func moveRedLeft[T constraints.Ordered](h *Node[T]) *Node[T] {
	h.colorFlip()
	if isRed(h.Right.Left) {
		h.Right = rotateRight(h.Right)
		h = rotateLeft(h)
		h.colorFlip()
	}
	return h
}

func moveRedRight[T constraints.Ordered](h *Node[T]) *Node[T] {
	h.colorFlip()
	if isRed(h.Left.Left) {
		h = rotateRight(h)
		h.colorFlip()
	}
	return h
}

func (h *Node[T]) fixUp() *Node[T] {
	if isRed(h.Right) { // fix right-learning reds on the way up
		h = rotateLeft(h)
	}
	if isRed(h.Left) && isRed(h.Left.Left) { // fix two reds in a row on the way up
		h = rotateRight(h)
	}
	if isRed(h.Left) && isRed(h.Right) { // split 4 nodes on the way up
		h.colorFlip()
	}
	return h
}

type Iterator[T constraints.Ordered] func(*Node[T]) bool

func (tree *rbTree[T]) preOrderPrint() string {
	arr := []T{}
	tree.PreOrderIterate(func(n *Node[T]) bool {
		arr = append(arr, n.Key)
		return true
	})
	return fmt.Sprintf("%v", arr)
}

func (tree *rbTree[T]) PreOrderIterate(iter Iterator[T]) {
	if tree.root != nil{
		tree.root.preOrderIterate(iter)
	}
}

func (n *Node[T]) preOrderIterate(iter Iterator[T]) (hit bool) {
	if !iter(n) {
		return false
	}
	if n.Left != nil && !n.Left.preOrderIterate(iter) {
		return false
	}
	if n.Right != nil && !n.Right.preOrderIterate(iter) {
		return false
	}
	return true
}
