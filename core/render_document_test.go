package core

import (
	"bytes"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestRenderDocument(t *testing.T) {
	tests := []struct {
		name     string
		children []ast.Node
		expected string
	}{
		{
			name: "Document with multiple paragraphs",
			children: []ast.Node{
				createParagraph("Paragraph 1"),
				createParagraph("Paragraph 2"),
			},
			expected: "Paragraph 1\n\nParagraph 2\n",
		},
		{
			name:     "Empty document",
			children: []ast.Node{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			doc := ast.NewDocument()
			for _, child := range tt.children {
				doc.AppendChild(doc, child)
			}
			renderDocument(&buf, doc, nil, 0)
			if buf.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, buf.String())
			}
		})
	}
}

func createParagraph(content string) *ast.Paragraph {
	p := ast.NewParagraph()
	p.AppendChild(p, ast.NewString([]byte(content)))
	return p
}
