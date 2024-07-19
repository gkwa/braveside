package core

import (
	"fmt"
	"io"

	"github.com/yuin/goldmark/ast"
)

func renderList(w io.Writer, n *ast.List, source []byte, level int) {
	start := n.Start
	if start == 0 {
		start = 1
	}
	for i, c := 0, n.FirstChild(); c != nil; i, c = i+1, c.NextSibling() {
		fmt.Fprint(w, indent(level))
		if n.IsOrdered() {
			fmt.Fprintf(w, "%d. ", start+i)
		} else {
			fmt.Fprint(w, "- ")
		}
		renderMarkdown(w, c, source, level+1)
		fmt.Fprintln(w)
	}
}

