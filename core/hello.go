package core

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/go-logr/logr"
	"github.com/yuin/goldmark/ast"
	"gopkg.in/yaml.v3"
)

type DefaultASTPrinter struct{}

func (p *DefaultASTPrinter) PrintAST(doc ast.Node, input []byte) {
	fmt.Println("AST structure:")
	printNode(doc, input, 0)
}

type DefaultFrontMatterProcessor struct{}

func (p *DefaultFrontMatterProcessor) ProcessFrontMatter(metaData map[string]interface{}) (string, error) {
	var frontMatterBuf bytes.Buffer
	encoder := yaml.NewEncoder(&frontMatterBuf)
	encoder.SetIndent(2)
	if err := encoder.Encode(metaData); err != nil {
		return "", err
	}
	encoder.Close()
	return frontMatterBuf.String(), nil
}

type DefaultMarkdownRenderer struct{}

func (r *DefaultMarkdownRenderer) RenderMarkdown(doc ast.Node, source []byte) ([]byte, error) {
	var contentBuf bytes.Buffer
	renderMarkdown(&contentBuf, doc, source, 0)
	return contentBuf.Bytes(), nil
}

func Hello(logger logr.Logger, showAST bool) error {
	input, err := os.ReadFile("testdata/input.md")
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	astPrinter := &DefaultASTPrinter{}
	frontMatterProcessor := &DefaultFrontMatterProcessor{}
	markdownRenderer := &DefaultMarkdownRenderer{}

	processor := NewMarkdownProcessor(astPrinter, frontMatterProcessor, markdownRenderer)
	output, err := processor.ProcessMarkdown(input)
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
