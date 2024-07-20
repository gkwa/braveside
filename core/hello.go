package core

import (
	"bytes"
	"context"
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

type contextKey int

const (
	loggerKey contextKey = iota
	showASTKey
)

func ContextWithLogger(ctx context.Context, logger logr.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func LoggerFrom(ctx context.Context) logr.Logger {
	logger, _ := ctx.Value(loggerKey).(logr.Logger)
	return logger
}

func ContextWithShowAST(ctx context.Context, showAST bool) context.Context {
	return context.WithValue(ctx, showASTKey, showAST)
}

func ShowASTFrom(ctx context.Context) bool {
	showAST, _ := ctx.Value(showASTKey).(bool)
	return showAST
}

func Hello(ctx context.Context) error {
	logger := LoggerFrom(ctx)

	input, err := os.ReadFile("testdata/input.md")
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	output, err := ProcessMarkdown(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to process markdown: %w", err)
	}

	err = os.WriteFile("output.md", output, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return compareDiff(logger, "testdata/input.md", "output.md")
}

func ProcessMarkdown(ctx context.Context, input []byte) ([]byte, error) {
	showAST := ShowASTFrom(ctx)

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

	var frontMatterBuf bytes.Buffer
	encoder := yaml.NewEncoder(&frontMatterBuf)
	encoder.SetIndent(2)
	if err := encoder.Encode(metaData); err != nil {
		return nil, err
	}
	encoder.Close()

	output := fmt.Sprintf("---\n%s---\n\n%s", frontMatterBuf.String(), contentBuf.String())
	return []byte(output), nil
}

func compareDiff(logger logr.Logger, file1, file2 string) error {
	cmd := exec.Command("diff", "--unified", "--ignore-blank-lines", "--ignore-all-space", file1, file2)
	diff, err := cmd.CombinedOutput()
	if err != nil && err.(*exec.ExitError).ExitCode() != 1 {
		return fmt.Errorf("diff command failed: %w", err)
	}

	if len(diff) > 0 {
		logger.Info("Differences found:", "diff", string(diff))
	} else {
		logger.Info("No differences found between input.md and output.md")
	}

	return nil
}
