package core

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestRenderParagraph(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:    "Simple paragraph",
			content: "This is a simple paragraph.",
			expected: `
This is a simple paragraph.

`,
		},
		{
			name:    "Paragraph with multiple lines",
			content: "This is a paragraph\nwith multiple lines.",
			expected: `
This is a paragraph
with multiple lines.

`,
		},
		{
			name:    "Paragraph with definition",
			content: "Term\n: Definition",
			expected: `
Term
: Definition

`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			para := ast.NewParagraph()
			para.AppendChild(para, ast.NewString([]byte(tt.content)))
			renderParagraph(&buf, para, []byte(tt.content), 0)
			if strings.TrimSpace(buf.String()) != strings.TrimSpace(tt.expected) {
				t.Errorf("Expected:\n%s\nGot:\n%s", strings.TrimSpace(tt.expected), strings.TrimSpace(buf.String()))
			}
		})
	}
}
