package rbTree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getIntTreeforTest() *rbTree[int] {
	return &rbTree[int]{
		root: &Node[int]{
			Key: 17,
			Left: &Node[int]{
				Key: 9,
				Left: &Node[int]{
					Key: 3,
				},
				Right: &Node[int]{
					Key: 12,
				},
			},
			Right: &Node[int]{
				Key: 19,
				Left: &Node[int]{
					Key: 18,
				},
				Right: &Node[int]{
					Key: 75,
					Left: &Node[int]{
						Key: 24,
					},
					Right: &Node[int]{
						Key: 81,
					},
				},
			},
		},
	}
}

func getStringTreeforTest() *rbTree[string] {
	return &rbTree[string]{
		root: &Node[string]{
			Key: "R",
			Left: &Node[string]{
				Key:   "E",
				Color: RED,
				Left: &Node[string]{
					Key: "C",
					Left: &Node[string]{
						Key:   "A",
						Color: RED,
					},
				},
				Right: &Node[string]{
					Key: "H",
				},
			},
			Right: &Node[string]{
				Key: "S",
			},
		},
	}

}

func TestHasOp(t *testing.T) {
	tree := getIntTreeforTest()
	keysExpToExist := []int{17, 9, 3, 12, 19, 18, 75, 24, 81}
	keysNotExpToExist := []int{-1, 0, 1}

	for _, ele := range keysExpToExist {
		assert.Truef(t, tree.Has(ele), "Element %d should exist in the tree", ele)
	}

	for _, ele := range keysNotExpToExist {
		assert.Falsef(t, tree.Has(ele), "Element %d should not exist in the tree", ele)
	}
}

func TestInsertOp(t *testing.T) {
	tree := NewTree[int]()
	for i := 0; i < 10; i++ {
		tree.InsertOrReplace(i)
	}
	assert.Equal(t, "[3 1 0 2 7 5 4 6 9 8]", tree.preOrderPrint())
}

func TestInsertOrReplaceWithString(t *testing.T) {
	act := []string{
		"[S]",
		"[S E]",
		"[E A S]",
		"[E A S R]",
		"[E C A S R]",
		"[R E C A H S]",
		"[R E C A H X S]",
		"[R E C A M H X S]",
		"[M E C A H R P X S]",
		"[M E C A L H R P X S]",
	}

	tree := NewTree[string]()
	elements := []string{"S", "E", "A", "R", "C", "H", "X", "M", "P", "L"}
	for i, ele := range elements {
		tree.InsertOrReplace(ele)
		assert.Equal(t, act[i], tree.preOrderPrint())
	}
}

func TestDeleteWithString(t *testing.T) {
	act := []string{
		"[]",
		"[S]",
		"[S E]",
		"[E A S]",
		"[R E A S]",
		"[R C A E S]",
		"[R E C A H S]",
		"[R E C A H X S]",
		"[M E C A H S R X]",
		"[M E C A H R P X S]",
	}
	tree := NewTree[string]()
	elements := []string{"S", "E", "A", "R", "C", "H", "X", "M", "P", "L"}
	for _, ele := range elements {
		tree.InsertOrReplace(ele)
	}

	for i := len(elements) - 1; i >= 0; i-- {
		tree.Delete(elements[i])
		assert.Equal(t, act[i], tree.preOrderPrint())
	}

}

func TestDeleteOp(t *testing.T) {
	tree := NewTree[int]()
	for i := 0; i < 3; i++ {
		tree.InsertOrReplace(i)
	}
	assert.Equal(t, "[1 0 2]", tree.preOrderPrint())
	for i := 0; i < 3; i++ {
		tree.Delete(i)
	}
	assert.Equal(t, "[]", tree.preOrderPrint())
}

func TestInOrderPrint(t *testing.T) {
	tree := getIntTreeforTest()
	assert.Equal(t, "[17 9 3 12 19 18 75 24 81]", tree.preOrderPrint())
}
