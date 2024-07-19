package core

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/go-logr/logr"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	east "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func Hello(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Hello function")
	logger.Info("Hello, World!")
	logger.V(1).Info("Debug: Exiting Hello function")

	input, err := os.ReadFile("testdata/input.md")
	if err != nil {
		log.Fatal(err)
	}

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

	doc := md.Parser().Parse(text.NewReader(input))

	fmt.Println("AST structure:")
	printNode(doc, input, 0)

	var buf bytes.Buffer
	renderMarkdown(&buf, doc, input)

	err = os.WriteFile("output.md", buf.Bytes(), 0o644)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("diff", "-u", "testdata/input.md", "output.md")
	diff, _ := cmd.CombinedOutput()
	if len(diff) > 0 {
		fmt.Println("Differences found:")
		fmt.Println(string(diff))
	} else {
		fmt.Println("No differences found between input.md and output.md")
	}
}

func indent(level int) string {
	return strings.Repeat("  ", level)
}

func printNode(n ast.Node, source []byte, level int) {
	fmt.Printf("%s%T\n", indent(level), n)
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		printNode(c, source, level+1)
	}
}

func renderMarkdown(w io.Writer, n ast.Node, source []byte) {
	switch v := n.(type) {
	case *ast.Document:
		for c := v.FirstChild(); c != nil; c = c.NextSibling() {
			renderMarkdown(w, c, source)
			if c.NextSibling() != nil {
				fmt.Fprintln(w)
			}
		}
	case *ast.Heading:
		for i := 0; i < 5; i++ {
			fmt.Fprintln(w)
		}

		fmt.Fprintf(w, "%s ", strings.Repeat("#", v.Level))
		for c := v.FirstChild(); c != nil; c = c.NextSibling() {
			renderMarkdown(w, c, source)
		}
		fmt.Fprintln(w)
	case *ast.Paragraph:
		for c := v.FirstChild(); c != nil; c = c.NextSibling() {
			renderMarkdown(w, c, source)
		}
		fmt.Fprintln(w)
	case *ast.Link:
		fmt.Fprintf(w, "[")
		for c := v.FirstChild(); c != nil; c = c.NextSibling() {
			renderMarkdown(w, c, source)
		}
		fmt.Fprintf(w, "](%s)", v.Destination)
	case *ast.Text:
		fmt.Fprint(w, string(v.Text(source)))
	case *ast.String:
		fmt.Fprint(w, string(v.Value))
	case *ast.AutoLink:
		fmt.Fprintf(w, "%s", v.URL(source))
	case *ast.List:
		for i, c := 0, v.FirstChild(); c != nil; i, c = i+1, c.NextSibling() {
			if v.IsOrdered() {
				fmt.Fprintf(w, "%d. ", i+1)
			} else {
				fmt.Fprint(w, "- ")
			}
			renderMarkdown(w, c, source)
			fmt.Fprintln(w)
		}
	case *ast.ListItem:
		for c := v.FirstChild(); c != nil; c = c.NextSibling() {
			renderMarkdown(w, c, source)
		}
	case *ast.Blockquote:
		fmt.Fprint(w, "> ")
		for c := v.FirstChild(); c != nil; c = c.NextSibling() {
			renderMarkdown(w, c, source)
		}
	case *ast.FencedCodeBlock:
		fmt.Fprint(w, "```")
		language := v.Language(source)
		if language != nil {
			fmt.Fprintf(w, "%s", language)
		}
		fmt.Fprintln(w)
		for i := 0; i < v.Lines().Len(); i++ {
			line := v.Lines().At(i)
			fmt.Fprintf(w, "%s", line.Value(source))
		}
		fmt.Fprintln(w, "```")
	case *east.Table:
		columnWidths := make([]int, 0)
		for r := v.FirstChild(); r != nil; r = r.NextSibling() {
			for i, c := 0, r.FirstChild(); c != nil; i, c = i+1, c.NextSibling() {
				cellLength := len(c.Text(source))
				if i >= len(columnWidths) {
					columnWidths = append(columnWidths, cellLength)
				} else if cellLength > columnWidths[i] {
					columnWidths[i] = cellLength
				}
			}
		}

		for r := v.FirstChild(); r != nil; r = r.NextSibling() {
			fmt.Fprint(w, "|")
			for i, c := 0, r.FirstChild(); c != nil; i, c = i+1, c.NextSibling() {
				fmt.Fprintf(w, " %-*s |", columnWidths[i], string(c.Text(source)))
			}
			fmt.Fprintln(w)
			if r == v.FirstChild() {
				fmt.Fprint(w, "|")
				for _, width := range columnWidths {
					fmt.Fprintf(w, "%s|", strings.Repeat("-", width+2))
				}
				fmt.Fprintln(w)
			}
		}
	case *east.TaskCheckBox:
		if v.IsChecked {
			fmt.Fprint(w, "[x] ")
		} else {
			fmt.Fprint(w, "[ ] ")
		}
	case *ast.Emphasis:
		marker := "*"
		if v.Level == 2 {
			marker = "**"
		}
		fmt.Fprint(w, marker)
		for c := v.FirstChild(); c != nil; c = c.NextSibling() {
			renderMarkdown(w, c, source)
		}
		fmt.Fprint(w, marker)
	case *ast.CodeSpan:
		fmt.Fprint(w, "`")
		for c := v.FirstChild(); c != nil; c = c.NextSibling() {
			renderMarkdown(w, c, source)
		}
		fmt.Fprint(w, "`")
	case *ast.Image:
		fmt.Fprintf(w, "![%s](%s)", v.Text(source), v.Destination)
	default:
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
}
