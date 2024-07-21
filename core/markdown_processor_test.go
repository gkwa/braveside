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
			name: "With frontmatter",
			input: `---
title: Test Document
author: John Doe
---

# Hello, World!

This is a test document.`,
			expected: `---
author: John Doe
title: Test Document
---

# Hello, World!

This is a test document.
`,
		},
		{
			name: "Without frontmatter",
			input: `# Hello, World!

This is a test document without frontmatter.`,
			expected: `# Hello, World!

This is a test document without frontmatter.
`,
		},
		{
			name: "With subscript",
			input: `# Chemical Formula

The chemical formula for water is H<sub>2</sub>O.`,
			expected: `# Chemical Formula

The chemical formula for water is H<sub>2</sub>O.
`,
		},
	}

	processor := NewMarkdownProcessor()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
