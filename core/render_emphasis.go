package core

import (
	"fmt"
	"io"

	"github.com/yuin/goldmark/ast"
)

func renderEmphasis(w io.Writer, n *ast.Emphasis, source []byte, level int) {
	marker := "*"
	if n.Level == 2 {
		marker = "**"
	}
	fmt.Fprint(w, marker)
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
	fmt.Fprint(w, marker)
}
