package core

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yuin/goldmark/ast"
	east "github.com/yuin/goldmark/extension/ast"
)

func TestRenderTable(t *testing.T) {
	tests := []struct {
		name     string
		rows     [][]string
		expected string
	}{
		{
			name: "Simple table",
			rows: [][]string{
				{"Header 1", "Header 2"},
				{"Row 1, Col 1", "Row 1, Col 2"},
				{"Row 2, Col 1", "Row 2, Col 2"},
			},
			expected: `
| Header 1     | Header 2     |
|--------------|--------------|
| Row 1, Col 1 | Row 1, Col 2 |
| Row 2, Col 1 | Row 2, Col 2 |
`,
		},
		{
			name: "Table with empty cells",
			rows: [][]string{
				{"Header 1", "Header 2", "Header 3"},
				{"Row 1, Col 1", "", "Row 1, Col 3"},
				{"", "Row 2, Col 2", ""},
			},
			expected: `
| Header 1     | Header 2     | Header 3     |
|--------------|--------------|--------------|
| Row 1, Col 1 |              | Row 1, Col 3 |
|              | Row 2, Col 2 |              |
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			table := east.NewTable()
			alignments := make([]east.Alignment, len(tt.rows[0]))
			for i := range alignments {
				alignments[i] = east.AlignNone
			}
			for _, row := range tt.rows {
				tr := east.NewTableRow(alignments)
				for _, cell := range row {
					tc := east.NewTableCell()
					tc.AppendChild(tc, ast.NewString([]byte(cell)))
					tr.AppendChild(tr, tc)
				}
				table.AppendChild(table, tr)
			}
			renderTable(&buf, table, []byte(strings.Join(tt.rows[0], "|")))
			if strings.TrimSpace(buf.String()) != strings.TrimSpace(tt.expected) {
				t.Errorf("Expected:\n%s\nGot:\n%s", strings.TrimSpace(tt.expected), strings.TrimSpace(buf.String()))
			}
		})
	}
}
