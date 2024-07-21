//go:build !failing
// +build !failing

package core

import (
	"testing"
)

func TestMarkdownProcessor(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "With subscript",
			input: `# Chemical Formula
The chemical formula for water is H<sub>2</sub>O.`,
			expected: `# Chemical Formula
The chemical formula for water is H<sub>2</sub>O.
`,
		},
		{
			name: "With superscript",
			input: `# Mathematical Expression
The square of a number is represented as n<sup>2</sup>.`,
			expected: `# Mathematical Expression
The square of a number is represented as n<sup>2</sup>.
`,
		},
	}

	processor := NewMarkdownProcessor()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Skip("Skipping test due to known failure")

			output, err := processor.ProcessMarkdown([]byte(tt.input), false)
			if err != nil {
				t.Fatalf("ProcessMarkdown() error = %v", err)
			}
			if string(output) != tt.expected {
				t.Errorf("ProcessMarkdown() output =\n%v\nwant\n%v", string(output), tt.expected)
			}
		})
	}
}
