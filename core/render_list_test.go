package core

import (
	"bytes"
	"regexp"
	"strings"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func TestRenderList(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "List with nested items followed by a term",
			input: `1. First item
   - Subitem
   - Another subitem
2. Second item

Term`,
			expected: `1\. First item
\s+- Subitem
\s+- Another subitem
2\. Second item

Term`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := goldmark.New(
				goldmark.WithExtensions(extension.GFM),
				goldmark.WithParserOptions(
					parser.WithAutoHeadingID(),
				),
			)

			doc := md.Parser().Parse(text.NewReader([]byte(tt.input)))

			var buf bytes.Buffer
			renderMarkdown(&buf, doc, []byte(tt.input), 0)

			output := strings.TrimSpace(buf.String())
			expectedRegex := regexp.MustCompile("(?m)^" + strings.TrimSpace(tt.expected) + "$")

			if !expectedRegex.MatchString(output) {
				t.Errorf("Output does not match expected pattern.\nExpected pattern:\n%s\nGot:\n%s", tt.expected, output)
			}
		})
	}
}
