package core

import (
	"fmt"
	"io"

	"github.com/yuin/goldmark/ast"
)

func renderFencedCodeBlock(w io.Writer, n *ast.FencedCodeBlock, source []byte) {
	fmt.Fprint(w, "```")
	language := n.Language(source)
	if language != nil {
		fmt.Fprintf(w, "%s", language)
	}
	fmt.Fprintln(w)
	for i := 0; i < n.Lines().Len(); i++ {
		line := n.Lines().At(i)
		fmt.Fprintf(w, "%s", line.Value(source))
	}
	fmt.Fprintln(w, "```")
}
