package core

import (
	"fmt"
	"io"

	"github.com/yuin/goldmark/ast"
)

func renderListItem(w io.Writer, n *ast.ListItem, source []byte, level int) {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		if c.Kind() == ast.KindList {
			fmt.Fprintln(w)
			renderList(w, c.(*ast.List), source, level)
		} else {
			renderMarkdown(w, c, source, level)
			if c.NextSibling() != nil && c.NextSibling().Kind() != ast.KindList {
				fmt.Fprint(w, " ")
			}
		}
	}
}
