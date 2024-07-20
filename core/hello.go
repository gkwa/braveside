package core

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/go-logr/logr"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"gopkg.in/yaml.v3"
)

func Hello(logger logr.Logger, showAST bool) {
	logger.V(1).Info("Debug: Entering Hello function")
	logger.Info("Hello, World!")
	logger.V(1).Info("Debug: Exiting Hello function")

	input, err := os.ReadFile("testdata/input.md")
	if err != nil {
		log.Fatal(err)
	}

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.DefinitionList, meta.Meta),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

	context := parser.NewContext()
	doc := md.Parser().Parse(text.NewReader(input), parser.WithContext(context))

	metaData := meta.Get(context)
	fmt.Printf("Frontmatter: %v\n", metaData)

	if showAST {
		fmt.Println("AST structure:")
		printNode(doc, input, 0)
	}

	var buf bytes.Buffer
	renderMarkdown(&buf, doc, input, 0)

	frontMatter, err := yaml.Marshal(metaData)
	if err != nil {
		log.Fatal(err)
	}

	output := fmt.Sprintf("---\n%s---\n\n%s", frontMatter, buf.String())

	err = os.WriteFile("output.md", []byte(output), 0o644)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("diff", "--unified", "--ignore-blank-lines", "--ignore-all-space", "testdata/input.md", "output.md")
	diff, _ := cmd.CombinedOutput()
	if len(diff) > 0 {
		fmt.Println("Differences found:")
		fmt.Println(string(diff))
	} else {
		fmt.Println("No differences found between input.md and output.md")
	}
}
