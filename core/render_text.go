package core

import (
	"fmt"
	"io"

	"github.com/yuin/goldmark/ast"
)

func renderText(w io.Writer, n *ast.Text, source []byte) {
	fmt.Fprint(w, string(n.Text(source)))
}
