package core

import (
	"bytes"
	"fmt"
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

func Hello(logger logr.Logger, showAST bool) error {
	input, err := os.ReadFile("testdata/input.md")
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	output, err := ProcessMarkdown(input, showAST)
	if err != nil {
		return fmt.Errorf("failed to process markdown: %w", err)
	}

	err = os.WriteFile("output.md", output, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	diff, err := compareDiff(logger, "testdata/input.md", "output.md")
	if err != nil {
		return fmt.Errorf("error comparing, %v", err)
	}

	fmt.Println(diff)
	return nil
}

func ProcessMarkdown(input []byte, showAST bool) ([]byte, error) {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.DefinitionList, meta.Meta),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

	context := parser.NewContext()
	doc := md.Parser().Parse(text.NewReader(input), parser.WithContext(context))

	metaData := meta.Get(context)

	if showAST {
		fmt.Println("AST structure:")
		printNode(doc, input, 0)
	}

	var contentBuf bytes.Buffer
	renderMarkdown(&contentBuf, doc, input, 0)

	var output []byte
	if len(metaData) == 0 {
		output = contentBuf.Bytes()
	} else {
		var frontMatterBuf bytes.Buffer
		encoder := yaml.NewEncoder(&frontMatterBuf)
		encoder.SetIndent(2)
		if err := encoder.Encode(metaData); err != nil {
			return nil, err
		}
		encoder.Close()

		output = []byte(fmt.Sprintf("---\n%s---\n\n%s", frontMatterBuf.String(), contentBuf.String()))
	}

	return output, nil
}

func compareDiff(logger logr.Logger, file1, file2 string) (string, error) {
	opts := []string{
		"--unified",
		"--ignore-blank-lines",
		"--ignore-all-space",
	}
	args := append(opts, file1, file2)
	cmd := exec.Command("diff", args...)
	diff, err := cmd.CombinedOutput()
	if err != nil && err.(*exec.ExitError).ExitCode() != 1 {
		return "", fmt.Errorf("diff command failed: %w", err)
	}

	var result string
	if len(diff) > 0 {
		result = fmt.Sprintf("Differences found:\ndiff%s", string(diff))
	} else {
		result = "No differences found between input.md and output.md"
		logger.Info(result)
	}

	return result, nil
}
