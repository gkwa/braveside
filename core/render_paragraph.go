package core

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/yuin/goldmark/ast"
)

func renderParagraph(w io.Writer, n *ast.Paragraph, source []byte, level int) {
	content := bytes.Buffer{}
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(&content, c, source, level)
	}
	paragraphContent := content.String()

	if strings.Contains(paragraphContent, "\n:") {
		parts := strings.SplitN(paragraphContent, "\n:", 2)
		fmt.Fprintf(w, "%s\n", strings.TrimSpace(parts[0]))
		fmt.Fprintf(w, ": %s", strings.TrimSpace(parts[1]))
	} else {
		fmt.Fprint(w, paragraphContent)
	}
	fmt.Fprintln(w)
}
