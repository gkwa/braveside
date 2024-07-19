package core

import (
	"fmt"
	"io"
)

func renderThematicBreak(w io.Writer) {
	fmt.Fprintln(w, "---")
}
