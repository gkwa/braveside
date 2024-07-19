package core

import (
	"bytes"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

func TestRenderMarkdown(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple paragraph",
			input:    "Hello, world!",
			expected: "Hello, world!\n",
		},
		{
			name:     "Heading",
			input:    "# Title",
			expected: "\n\n\n\n\n# Title\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			md := goldmark.New()
			doc := md.Parser().Parse(text.NewReader([]byte(tt.input)))

			for node := doc.FirstChild(); node != nil; node = node.NextSibling() {
				renderMarkdown(&buf, node, []byte(tt.input), 0)
			}

			if buf.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, buf.String())
			}
		})
	}
}
