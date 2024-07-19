package core

import (
	"bytes"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestRenderString(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "Simple string",
			content:  "Hello, world!",
			expected: "Hello, world!",
		},
		{
			name:     "Empty string",
			content:  "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			str := ast.NewString([]byte(tt.content))
			renderString(&buf, str)
			if buf.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, buf.String())
			}
		})
	}
}
