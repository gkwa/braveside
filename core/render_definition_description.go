package core

import (
	"fmt"
	"io"

	east "github.com/yuin/goldmark/extension/ast"
)

func renderDefinitionDescription(w io.Writer, n *east.DefinitionDescription, source []byte, level int) {
	fmt.Fprint(w, ": ")
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
}
