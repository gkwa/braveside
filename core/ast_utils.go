package core

import (
	"fmt"

	"github.com/yuin/goldmark/ast"
)

func printNode(n ast.Node, source []byte, level int) {
	fmt.Printf("%s%T\n", indent(level), n)
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		printNode(c, source, level+1)
	}
}
