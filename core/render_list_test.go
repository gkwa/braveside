package core

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestRenderList(t *testing.T) {
	tests := []struct {
		name      string
		isOrdered bool
		items     []string
		expected  string
	}{
		{
			name:      "Unordered list",
			isOrdered: false,
			items:     []string{"Item 1", "Item 2", "Item 3"},
			expected: `
- Item 1
- Item 2
- Item 3
`,
		},
		{
			name:      "Ordered list",
			isOrdered: true,
			items:     []string{"First", "Second", "Third"},
			expected: `
1. First
2. Second
3. Third
`,
		},
		{
			name:      "Empty list",
			isOrdered: false,
			items:     []string{},
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			var marker byte
			if tt.isOrdered {
				marker = '.'
			} else {
				marker = '-'
			}
			list := ast.NewList(marker)
			for _, item := range tt.items {
				li := ast.NewListItem(0)
				li.AppendChild(li, ast.NewString([]byte(item)))
				list.AppendChild(list, li)
			}
			renderList(&buf, list, nil, 0)
			if strings.TrimSpace(buf.String()) != strings.TrimSpace(tt.expected) {
				t.Errorf("Expected:\n%s\nGot:\n%s", strings.TrimSpace(tt.expected), strings.TrimSpace(buf.String()))
			}
		})
	}
}
