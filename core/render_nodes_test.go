package core

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yuin/goldmark/ast"
	east "github.com/yuin/goldmark/extension/ast"
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

func TestRenderEmphasis(t *testing.T) {
	tests := []struct {
		name     string
		level    int
		content  string
		expected string
	}{
		{
			name:     "Italic emphasis",
			level:    1,
			content:  "italic text",
			expected: "*italic text*",
		},
		{
			name:     "Bold emphasis",
			level:    2,
			content:  "bold text",
			expected: "**bold text**",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			em := ast.NewEmphasis(tt.level)
			em.AppendChild(em, ast.NewString([]byte(tt.content)))
			renderEmphasis(&buf, em, nil, 0)
			if buf.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, buf.String())
			}
		})
	}
}

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
			expected: "# Main Title\n",
		},
		{
			name:     "H3 heading",
			level:    3,
			content:  "Subheading",
			expected: "### Subheading\n",
		},
		{
			name:     "H6 heading",
			level:    6,
			content:  "Lowest level",
			expected: "###### Lowest level\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			heading := ast.NewHeading(tt.level)
			heading.AppendChild(heading, ast.NewString([]byte(tt.content)))
			renderHeading(&buf, heading, []byte(tt.content), 0)
			if buf.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, buf.String())
			}
		})
	}
}

func TestRenderImage(t *testing.T) {
	tests := []struct {
		name     string
		alt      string
		src      string
		expected string
	}{
		{
			name:     "Basic image",
			alt:      "Alt text",
			src:      "image.jpg",
			expected: "![Alt text](image.jpg)",
		},
		{
			name:     "Image with no alt text",
			alt:      "",
			src:      "image.png",
			expected: "![](image.png)",
		},
		{
			name:     "Image with special characters",
			alt:      "Image & Text",
			src:      "https://example.com/image.gif?param=value",
			expected: "![Image & Text](https://example.com/image.gif?param=value)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			node := ast.NewImage(ast.NewLink())
			node.Destination = []byte(tt.src)
			if tt.alt != "" {
				node.AppendChild(node, ast.NewString([]byte(tt.alt)))
			}
			renderImage(&buf, node, nil)
			if buf.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, buf.String())
			}
		})
	}
}

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

