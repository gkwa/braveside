package core

import (
	"fmt"
	"io"

	"github.com/yuin/goldmark/ast"
)

func renderListItem(w io.Writer, n *ast.ListItem, source []byte, level int) {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
		if c.NextSibling() != nil {
			fmt.Fprintln(w)
		}
	}
}
