package core

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestRenderHeading(t *testing.T) {
	tests := []struct {
		name     string
		level    int
		content  string
		expected string
	}{
		{
			name:     "H1 heading",
			level:    1,
			content:  "Main Title",
			expected: "\n\n\n\n\n# Main Title\n",
		},
		{
			name:     "H3 heading",
			level:    3,
			content:  "Subheading",
			expected: "\n\n\n\n\n### Subheading\n",
		},
		{
			name:     "H6 heading",
			level:    6,
			content:  "Lowest level",
			expected: "\n\n\n\n\n###### Lowest level\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			heading := ast.NewHeading(tt.level)
			heading.AppendChild(heading, ast.NewString([]byte(tt.content)))
			renderHeading(&buf, heading, []byte(tt.content), 0)
			if strings.TrimSpace(buf.String()) != strings.TrimSpace(tt.expected) {
				t.Errorf("Expected %q, got %q", strings.TrimSpace(tt.expected), strings.TrimSpace(buf.String()))
			}
		})
	}
}
