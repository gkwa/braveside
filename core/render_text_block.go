package core

import (
	"io"

	"github.com/yuin/goldmark/ast"
)

func renderTextBlock(w io.Writer, n *ast.TextBlock, source []byte, level int) {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
}
