package core

import (
	"fmt"
	"io"

	"github.com/yuin/goldmark/ast"
)

func renderDefault(w io.Writer, n ast.Node, source []byte) {
	if n.Type() == ast.TypeBlock {
		for i := 0; i < n.Lines().Len(); i++ {
			line := n.Lines().At(i)
			fmt.Fprint(w, string(line.Value(source)))
			if i < n.Lines().Len()-1 {
				fmt.Fprintln(w)
			}
		}
	}
}
