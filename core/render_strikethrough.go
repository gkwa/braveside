package core

import (
	"fmt"
	"io"

	east "github.com/yuin/goldmark/extension/ast"
)

func renderStrikethrough(w io.Writer, n *east.Strikethrough, source []byte, level int) {
	fmt.Fprint(w, "~~")
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
	fmt.Fprint(w, "~~")
}
