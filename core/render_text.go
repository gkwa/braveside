package core

import (
	"fmt"
	"io"
	"os"

	"github.com/yuin/goldmark/ast"
)

func renderText(w io.Writer, n *ast.Text, source []byte) {
	text := n.Text(source)
	if n.SoftLineBreak() {
		text = append(text, '\n')
	}
	_, err := w.Write(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing text: %v\n", err)
	}
}
