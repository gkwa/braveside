package core

import (
	"io"

	east "github.com/yuin/goldmark/extension/ast"
)

func renderDefinitionTerm(w io.Writer, n *east.DefinitionTerm, source []byte, level int) {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
}
