package core

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestRenderList(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "List with nested items followed by a term",
			input: `
1. First item
   - Subitem
   - Another subitem
2. Second item

Term
`,
			expected: `
1. First item
   - Subitem
   - Another subitem
2. Second item

Term
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			doc := ast.NewDocument()
			list := ast.NewList('.')

			lines := strings.Split(tt.input, "\n")
			for i, line := range lines {
				if strings.HasPrefix(line, "Term") {
					para := ast.NewParagraph()
					para.AppendChild(para, ast.NewString([]byte("Term")))
					doc.AppendChild(doc, para)
					break
				}

				if strings.HasPrefix(line, "1.") || strings.HasPrefix(line, "2.") {
					item := ast.NewListItem(0)
					itemContent := strings.TrimSpace(strings.TrimPrefix(line, "1."))
					itemContent = strings.TrimSpace(strings.TrimPrefix(itemContent, "2."))
					item.AppendChild(item, ast.NewString([]byte(itemContent)))
					list.AppendChild(list, item)
				} else if strings.HasPrefix(line, "   -") {
					subItem := ast.NewListItem(0)
					subItemContent := strings.TrimSpace(strings.TrimPrefix(line, "   -"))
					subItem.AppendChild(subItem, ast.NewString([]byte(subItemContent)))
					list.LastChild().AppendChild(list.LastChild(), subItem)
				}

				if i == len(lines)-1 || strings.HasPrefix(lines[i+1], "Term") {
					doc.AppendChild(doc, list)
				}
			}

			renderMarkdown(&buf, doc, []byte(tt.input), 0)

			if buf.String() != tt.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s", tt.expected, buf.String())
			}
		})
	}
}
