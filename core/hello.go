package core

import (
	"context"
	"fmt"
	"os"

	"github.com/go-logr/logr"
)

var ShowAST bool

func Hello(ctx context.Context) error {
	logger := logr.FromContextOrDiscard(ctx)

	input, err := os.ReadFile("testdata/input.md")
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	frontMatterProcessor := &DefaultFrontMatterProcessor{}
	markdownRenderer := &DefaultMarkdownRenderer{}
	var processor *MarkdownProcessor

	if ShowAST {
		astPrinter := &DefaultASTPrinter{}
		processor = NewMarkdownProcessorWithASTPrinter(astPrinter, frontMatterProcessor, markdownRenderer)
	} else {
		processor = NewMarkdownProcessor(frontMatterProcessor, markdownRenderer)
	}

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
