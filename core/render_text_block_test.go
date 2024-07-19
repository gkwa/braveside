package core

import (
	"bytes"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestRenderTextBlock(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "Simple text block",
			content:  "This is a text block.",
			expected: "This is a text block.",
		},
		{
			name:     "Multi-line text block",
			content:  "Line 1\nLine 2\nLine 3",
			expected: "Line 1\nLine 2\nLine 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			block := ast.NewTextBlock()
			block.AppendChild(block, ast.NewString([]byte(tt.content)))
			renderTextBlock(&buf, block, []byte(tt.content), 0)
			if buf.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, buf.String())
			}
		})
	}
}
