package core

import (
	"fmt"
	"io"

	"github.com/yuin/goldmark/ast"
)

func renderImage(w io.Writer, n *ast.Image, source []byte) {
	fmt.Fprintf(w, "![%s](%s)", n.Text(source), n.Destination)
}
