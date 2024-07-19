package core

import (
	"fmt"
	"io"

	"github.com/yuin/goldmark/ast"
)

func renderString(w io.Writer, n *ast.String) {
	fmt.Fprint(w, string(n.Value))
}
