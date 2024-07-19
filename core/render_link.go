package core

import (
	"fmt"
	"io"

	"github.com/yuin/goldmark/ast"
)

func renderLink(w io.Writer, n *ast.Link, source []byte, level int) {
	fmt.Fprintf(w, "[")
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
	fmt.Fprintf(w, "](%s)", n.Destination)
}
