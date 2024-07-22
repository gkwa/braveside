package core

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-logr/logr"
)

func TestHello(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdata")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	inputPath := filepath.Join(tempDir, "input.md")
	inputContent := []byte(`# Test Input

This is a test input file.`)
	err = os.WriteFile(inputPath, inputContent, 0o644)
	if err != nil {
		t.Fatalf("Failed to create test input file: %v", err)
	}

	oldWd, _ := os.Getwd()
	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("Failed to change working directory: %v", err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Errorf("Failed to change back to original working directory: %v", err)
		}
	}()

	err = os.Mkdir("testdata", 0o755)
	if err != nil {
		t.Fatalf("Failed to create testdata directory: %v", err)
	}

	err = os.Rename(inputPath, "testdata/input.md")
	if err != nil {
		t.Fatalf("Failed to move input file: %v", err)
	}
}

type loggerKey struct{}

func LoggerFrom(ctx context.Context) logr.Logger {
	logger, _ := ctx.Value(loggerKey{}).(logr.Logger)
	return logger
}
