package core

import (
	"fmt"
	"io"

	"github.com/yuin/goldmark/ast"
)

func renderBlockquote(w io.Writer, n *ast.Blockquote, source []byte, level int) {
	fmt.Fprint(w, "> ")
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
}
