package core

import (
	"bytes"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestRenderLink(t *testing.T) {
	tests := []struct {
		name        string
		text        string
		destination string
		expected    string
	}{
		{
			name:        "Simple link",
			text:        "OpenAI",
			destination: "https://openai.com",
			expected:    "[OpenAI](https://openai.com)",
		},
		{
			name:        "Link with spaces in text",
			text:        "Google Search",
			destination: "https://google.com",
			expected:    "[Google Search](https://google.com)",
		},
		{
			name:        "Link with query parameters",
			text:        "Search",
			destination: "https://example.com/search?q=test&lang=en",
			expected:    "[Search](https://example.com/search?q=test&lang=en)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			link := ast.NewLink()
			link.AppendChild(link, ast.NewString([]byte(tt.text)))
			link.Destination = []byte(tt.destination)
			renderLink(&buf, link, []byte(tt.text), 0)
			if buf.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, buf.String())
			}
		})
	}
}
