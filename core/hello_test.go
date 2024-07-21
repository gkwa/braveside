package core

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/go-logr/zapr"
	"go.uber.org/zap/zaptest"
)

func TestHello(t *testing.T) {
	zapLogger := zaptest.NewLogger(t)
	logger := zapr.NewLogger(zapLogger)

	tempDir, err := os.MkdirTemp("", "testdata")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	inputPath := tempDir + "/input.md"
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

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err = Hello(logger, false)
	if err != nil {
		t.Fatalf("Hello() error = %v", err)
	}

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read captured output: %v", err)
	}
	output := buf.String()

	if !strings.Contains(output, "No differences found between input.md and output.md") {
		t.Errorf("Expected 'No differences found' message, got: %s", output)
	}

	if strings.Contains(output, "AST structure:") {
		t.Errorf("AST structure should not be printed when showAST is false")
	}

	r, w, _ = os.Pipe()
	os.Stdout = w

	err = Hello(logger, true)
	if err != nil {
		t.Fatalf("Hello() error = %v", err)
	}

	w.Close()
	os.Stdout = oldStdout

	buf.Reset()
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read captured output: %v", err)
	}
	output = buf.String()

	if !strings.Contains(output, "AST structure:") {
		t.Errorf("AST structure should be printed when showAST is true")
	}
}
