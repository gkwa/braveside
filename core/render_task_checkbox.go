package core

import (
	"fmt"
	"io"

	east "github.com/yuin/goldmark/extension/ast"
)

func renderTaskCheckBox(w io.Writer, n *east.TaskCheckBox) {
	if n.IsChecked {
		fmt.Fprint(w, "[x] ")
	} else {
		fmt.Fprint(w, "[ ] ")
	}
}
