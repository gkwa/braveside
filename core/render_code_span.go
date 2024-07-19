package core

import (
	"fmt"
	"io"

	"github.com/yuin/goldmark/ast"
)

func renderCodeSpan(w io.Writer, n *ast.CodeSpan, source []byte, level int) {
	fmt.Fprint(w, "`")
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
	fmt.Fprint(w, "`")
}
