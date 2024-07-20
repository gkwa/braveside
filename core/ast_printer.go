package core

import (
	"fmt"

	"github.com/yuin/goldmark/ast"
)

func printNode(n ast.Node, source []byte, level int) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}

	fmt.Printf("%s%s: %s\n", indent, n.Kind().String(), string(n.Text(source)))

	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		printNode(c, source, level+1)
	}
}
