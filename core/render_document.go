package core

import (
	"fmt"
	"io"

	"github.com/yuin/goldmark/ast"
)

func renderDocument(w io.Writer, n *ast.Document, source []byte, level int) {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
		if c.NextSibling() != nil {
			fmt.Fprintln(w)
		}
	}
}
