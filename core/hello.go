package core

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/go-logr/logr"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
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
		goldmark.WithExtensions(extension.GFM, extension.DefinitionList),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

	doc := md.Parser().Parse(text.NewReader(input))

	fmt.Println("AST structure:")
	printNode(doc, input, 0)

	var buf bytes.Buffer
	renderMarkdown(&buf, doc, input, 0)

	err = os.WriteFile("output.md", buf.Bytes(), 0o644)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("diff", "--unified", "--ignore-all-space", "testdata/input.md", "output.md")
	diff, _ := cmd.CombinedOutput()
	if len(diff) > 0 {
		fmt.Println("Differences found:")
		fmt.Println(string(diff))
	} else {
		fmt.Println("No differences found between input.md and output.md")
	}
}
