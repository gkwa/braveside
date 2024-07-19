package core

import (
	"fmt"
	"io"
	"strings"

	"github.com/yuin/goldmark/ast"
)

func renderHeading(w io.Writer, n *ast.Heading, source []byte, level int) {
	for i := 0; i < 5; i++ {
		fmt.Fprintln(w)
	}
	fmt.Fprintf(w, "%s ", strings.Repeat("#", n.Level))
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
	fmt.Fprintln(w)
}
