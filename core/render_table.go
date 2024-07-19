package core

import (
	"fmt"
	"io"
	"strings"

	east "github.com/yuin/goldmark/extension/ast"
)

func renderTable(w io.Writer, n *east.Table, source []byte) {
	columnWidths := make([]int, 0)
	for r := n.FirstChild(); r != nil; r = r.NextSibling() {
		for i, c := 0, r.FirstChild(); c != nil; i, c = i+1, c.NextSibling() {
			cellLength := len(strings.TrimSpace(string(c.Text(source))))
			if i >= len(columnWidths) {
				columnWidths = append(columnWidths, cellLength)
			} else if cellLength > columnWidths[i] {
				columnWidths[i] = cellLength
			}
		}
	}

	for r := n.FirstChild(); r != nil; r = r.NextSibling() {
		fmt.Fprint(w, "|")
		for i, c := 0, r.FirstChild(); c != nil; i, c = i+1, c.NextSibling() {
			cellContent := strings.TrimSpace(string(c.Text(source)))
			fmt.Fprintf(w, " %-*s |", columnWidths[i], cellContent)
		}
		fmt.Fprintln(w)
		if r == n.FirstChild() {
			fmt.Fprint(w, "|")
			for _, width := range columnWidths {
				fmt.Fprintf(w, "%s|", strings.Repeat("-", width+2))
			}
			fmt.Fprintln(w)
		}
	}
}
