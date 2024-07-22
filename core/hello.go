package core

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/go-logr/logr"
)

func ProcessInputMarkdown(ctx context.Context, input io.Reader, output io.Writer) error {
	inputBytes, err := io.ReadAll(input)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	frontMatterProcessor := &DefaultFrontMatterProcessor{}
	markdownRenderer := &DefaultMarkdownRenderer{}

	processor := NewMarkdownProcessor(frontMatterProcessor, markdownRenderer)
	outputBytes, err := processor.ProcessMarkdown(inputBytes)
	if err != nil {
		return fmt.Errorf("failed to process markdown: %w", err)
	}

	_, err = output.Write(outputBytes)
	if err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	return nil
}

func Hello(ctx context.Context, input io.Reader, output io.Writer, diffOutput io.Writer, compareDiffFunc func(logr.Logger, io.Reader, io.Reader) (string, error)) error {
	logger := logr.FromContextOrDiscard(ctx)

	var processedOutput bytes.Buffer
	err := ProcessInputMarkdown(ctx, input, &processedOutput)
	if err != nil {
		return err
	}

	_, err = io.Copy(output, &processedOutput)
	if err != nil {
		return fmt.Errorf("failed to copy processed output: %w", err)
	}

	diff, err := compareDiffFunc(logger, input, bytes.NewReader(processedOutput.Bytes()))
	if err != nil {
		return fmt.Errorf("error comparing, %v", err)
	}

	_, err = fmt.Fprintln(diffOutput, diff)
	return err
}
