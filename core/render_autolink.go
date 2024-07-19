package core

import (
	"fmt"
	"io"

	"github.com/yuin/goldmark/ast"
)

func renderAutoLink(w io.Writer, n *ast.AutoLink, source []byte) {
	fmt.Fprintf(w, "%s", n.URL(source))
}
