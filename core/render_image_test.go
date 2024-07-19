package core

import (
	"bytes"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestRenderImage(t *testing.T) {
	tests := []struct {
		name     string
		alt      string
		src      string
		expected string
	}{
		{
			name:     "Basic image",
			alt:      "Alt text",
			src:      "image.jpg",
			expected: "![Alt text](image.jpg)",
		},
		{
			name:     "Image with no alt text",
			alt:      "",
			src:      "image.png",
			expected: "![](image.png)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			node := ast.NewImage(ast.NewLink())
			node.Destination = []byte(tt.src)
			if tt.alt != "" {
				node.AppendChild(node, ast.NewString([]byte(tt.alt)))
			}
			renderImage(&buf, node, nil)
			if buf.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, buf.String())
			}
		})
	}
}
