package core

import (
	"fmt"
	"io"

	east "github.com/yuin/goldmark/extension/ast"
)

func renderDefinitionList(w io.Writer, n *east.DefinitionList, source []byte, level int) {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
		if c.NextSibling() != nil {
			fmt.Fprintln(w)
		}
	}
	fmt.Fprintln(w)
}
