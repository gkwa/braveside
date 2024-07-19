package core

import (
	"bytes"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestRenderCodeSpan(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "Simple code span",
			content:  "var x = 5",
			expected: "`var x = 5`",
		},
		{
			name:     "Code span with spaces",
			content:  "function hello() { }",
			expected: "`function hello() { }`",
		},
		{
			name:     "Empty code span",
			content:  "",
			expected: "``",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			codeSpan := ast.NewCodeSpan()
			codeSpan.AppendChild(codeSpan, ast.NewString([]byte(tt.content)))
			renderCodeSpan(&buf, codeSpan, []byte(tt.content), 0)
			if buf.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, buf.String())
			}
		})
	}
}
