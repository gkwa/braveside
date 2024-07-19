package core

import (
	"bytes"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestRenderEmphasis(t *testing.T) {
	tests := []struct {
		name     string
		level    int
		content  string
		expected string
	}{
		{
			name:     "Italic emphasis",
			level:    1,
			content:  "italic text",
			expected: "*italic text*",
		},
		{
			name:     "Bold emphasis",
			level:    2,
			content:  "bold text",
			expected: "**bold text**",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			em := ast.NewEmphasis(tt.level)
			em.AppendChild(em, ast.NewString([]byte(tt.content)))
			renderEmphasis(&buf, em, nil, 0)
			if buf.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, buf.String())
			}
		})
	}
}
