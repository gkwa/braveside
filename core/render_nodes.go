package core

import (
	"fmt"
	"io"
	"strings"

	"github.com/yuin/goldmark/ast"
	east "github.com/yuin/goldmark/extension/ast"
)

func renderAutoLink(w io.Writer, n *ast.AutoLink, source []byte) {
	fmt.Fprintf(w, "%s", n.URL(source))
}

func renderBlockquote(w io.Writer, n *ast.Blockquote, source []byte, level int) {
	fmt.Fprint(w, "> ")
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
}

func renderCodeSpan(w io.Writer, n *ast.CodeSpan, source []byte, level int) {
	fmt.Fprint(w, "`")
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
	fmt.Fprint(w, "`")
}

func renderDefinitionDescription(w io.Writer, n *east.DefinitionDescription, source []byte, level int) {
	fmt.Fprint(w, ": ")
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
}

func renderDefinitionList(w io.Writer, n *east.DefinitionList, source []byte, level int) {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
		if c.NextSibling() != nil {
			fmt.Fprintln(w)
		}
	}
	fmt.Fprintln(w)
}

func renderDefinitionTerm(w io.Writer, n *east.DefinitionTerm, source []byte, level int) {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
}

func renderDocument(w io.Writer, n *ast.Document, source []byte, level int) {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
		if c.NextSibling() != nil {
			fmt.Fprintln(w)
		}
	}
}

func renderEmphasis(w io.Writer, n *ast.Emphasis, source []byte, level int) {
	marker := "*"
	if n.Level == 2 {
		marker = "**"
	}
	fmt.Fprint(w, marker)
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
	fmt.Fprint(w, marker)
}

func renderFencedCodeBlock(w io.Writer, n *ast.FencedCodeBlock, source []byte) {
	fmt.Fprint(w, "```")
	language := n.Language(source)
	if language != nil {
		fmt.Fprintf(w, "%s", language)
	}
	fmt.Fprintln(w)
	for i := 0; i < n.Lines().Len(); i++ {
		line := n.Lines().At(i)
		fmt.Fprintf(w, "%s", line.Value(source))
	}
	fmt.Fprintln(w, "```")
}

func renderHeading(w io.Writer, n *ast.Heading, source []byte, level int) {
	fmt.Fprintf(w, "%s ", strings.Repeat("#", n.Level))
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
	fmt.Fprintln(w)
}

func renderImage(w io.Writer, n *ast.Image, source []byte) {
	fmt.Fprintf(w, "![%s](%s)", n.Text(source), n.Destination)
}

func renderLink(w io.Writer, n *ast.Link, source []byte, level int) {
	fmt.Fprintf(w, "[")
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
	fmt.Fprintf(w, "](%s)", n.Destination)
}

func renderList(w io.Writer, n *ast.List, source []byte, level int) {
	start := n.Start
	if start == 0 {
		start = 1
	}
	for i, c := 0, n.FirstChild(); c != nil; i, c = i+1, c.NextSibling() {
		fmt.Fprint(w, indent(level))
		if n.IsOrdered() {
			fmt.Fprintf(w, "%d. ", start+i)
		} else {
			fmt.Fprint(w, "- ")
		}
		renderListItem(w, c.(*ast.ListItem), source, level+1)
		if c.NextSibling() != nil {
			fmt.Fprintln(w)
		}
	}
	if n.Parent().Kind() != ast.KindListItem {
		fmt.Fprintln(w)
	}
}

func renderListItem(w io.Writer, n *ast.ListItem, source []byte, level int) {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		if c.Kind() == ast.KindList {
			fmt.Fprintln(w)
			renderList(w, c.(*ast.List), source, level)
		} else {
			renderMarkdown(w, c, source, level)
			if c.NextSibling() != nil && c.NextSibling().Kind() != ast.KindList {
				fmt.Fprint(w, " ")
			}
		}
	}
}

func renderParagraph(w io.Writer, n *ast.Paragraph, source []byte, level int) {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
	fmt.Fprintln(w)
}

func renderString(w io.Writer, n *ast.String) {
	fmt.Fprint(w, string(n.Value))
}

func renderStrikethrough(w io.Writer, n *east.Strikethrough, source []byte, level int) {
	fmt.Fprint(w, "~~")
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
	fmt.Fprint(w, "~~")
}

func renderTable(w io.Writer, n *east.Table, source []byte) {
	columnWidths := make([]int, 0)
	for r := n.FirstChild(); r != nil; r = r.NextSibling() {
		for i, c := 0, r.FirstChild(); c != nil; i, c = i+1, c.NextSibling() {
			cellLength := len(strings.TrimSpace(string(c.Text(source))))
			if i >= len(columnWidths) {
				columnWidths = append(columnWidths, cellLength)
			} else if cellLength > columnWidths[i] {
				columnWidths[i] = cellLength
			}
		}
	}

	for r := n.FirstChild(); r != nil; r = r.NextSibling() {
		fmt.Fprint(w, "|")
		for i, c := 0, r.FirstChild(); c != nil; i, c = i+1, c.NextSibling() {
			cellContent := strings.TrimSpace(string(c.Text(source)))
			fmt.Fprintf(w, " %-*s |", columnWidths[i], cellContent)
		}
		fmt.Fprintln(w)
		if r == n.FirstChild() {
			fmt.Fprint(w, "|")
			for _, width := range columnWidths {
				fmt.Fprintf(w, "%s|", strings.Repeat("-", width+2))
			}
			fmt.Fprintln(w)
		}
	}
}

func renderTaskCheckBox(w io.Writer, n *east.TaskCheckBox) {
	if n.IsChecked {
		fmt.Fprint(w, "[x] ")
	} else {
		fmt.Fprint(w, "[ ] ")
	}
}

func renderText(w io.Writer, n *ast.Text, source []byte) {
	fmt.Fprint(w, string(n.Text(source)))
}

func renderTextBlock(w io.Writer, n *ast.TextBlock, source []byte, level int) {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		renderMarkdown(w, c, source, level)
	}
}

func renderThematicBreak(w io.Writer) {
	fmt.Fprintln(w, "---")
}

func renderDefault(w io.Writer, n ast.Node, source []byte) {
	if n.Type() == ast.TypeBlock {
		for i := 0; i < n.Lines().Len(); i++ {
			line := n.Lines().At(i)
			fmt.Fprint(w, string(line.Value(source)))
			if i < n.Lines().Len()-1 {
				fmt.Fprintln(w)
			}
		}
	}
}
