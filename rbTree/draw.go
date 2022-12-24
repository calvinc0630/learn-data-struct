package rbTree

import (
	"fmt"
	"io"
	"text/template"

	"golang.org/x/exp/constraints"
)

var (
	graphTemplate = `strict graph {
{{range .Edges}}{{.From}} -- {{.To}} {{if .Color}}[color={{.Color}},penwidth=3]{{end}}
{{end}}
}
`
	execGraphTemplate = template.Must(template.New("graph").Parse(graphTemplate))
)

type edge struct {
	From  string
	To    string
	Color string
}

type graph struct {
	Edges []edge
}

func drawDot[T constraints.Ordered](tree *rbTree[T], w io.Writer) {
	g := graph{}
	tree.PreOrderIterate(func(n *Node[T]) bool {
		if n.Left != nil {
			edge := edge{
				From: fmt.Sprint(n.Key),
				To:   fmt.Sprint(n.Left.Key),
			}
			if n.Left.Color == RED {
				edge.Color = "RED"
			}
			g.Edges = append(g.Edges, edge)
		}
		if n.Right != nil {
			edge := edge{
				From: fmt.Sprint(n.Key),
				To:   fmt.Sprint(n.Right.Key),
			}
			if n.Right.Color == RED {
				edge.Color = "RED"
			}
			g.Edges = append(g.Edges, edge)
		}
		return true
	})
	execGraphTemplate.Execute(w, g)
}
