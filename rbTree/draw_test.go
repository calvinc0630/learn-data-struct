package rbTree

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestDrawDotOfIntTree(t *testing.T)(){
	tree := getIntTreeforTest()
	buf := bytes.Buffer{}
	drawDot(tree, &buf)

	assert.Equal(t, `strict graph {
17 -- 9
17 -- 19
9 -- 3
9 -- 12
19 -- 18
19 -- 75
75 -- 24
75 -- 81

}
`, strings.ReplaceAll(buf.String(), " \n", "\n"))
}

func TestDrawDotOfStringTree(t *testing.T)(){
	tree := getStringTreeforTest()
	buf := bytes.Buffer{}
	drawDot(tree, &buf)

	assert.Equal(t, `strict graph {
R -- E [color=RED,penwidth=3]
R -- S
E -- C
E -- H
C -- A [color=RED,penwidth=3]

}
`, strings.ReplaceAll(buf.String(), " \n", "\n"))
}