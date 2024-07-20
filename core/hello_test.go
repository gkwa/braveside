package core

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/go-logr/zapr"
	"go.uber.org/zap/zaptest"
)

func TestProcessMarkdown(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "With frontmatter",
			input: `---
title: Test Document
author: John Doe
---

# Hello, World!

This is a test document.`,
			expected: `---
author: John Doe
title: Test Document
---

# Hello, World!

This is a test document.
`,
		},
		{
			name: "Without frontmatter",
			input: `# Hello, World!

This is a test document without frontmatter.`,
			expected: `# Hello, World!

This is a test document without frontmatter.
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := ProcessMarkdown([]byte(tt.input), false)
			if err != nil {
				t.Fatalf("ProcessMarkdown() error = %v", err)
			}
			if string(output) != tt.expected {
				t.Errorf("ProcessMarkdown() output =\n%v\nwant\n%v", string(output), tt.expected)
			}
		})
	}
}

func TestHello(t *testing.T) {
	zapLogger := zaptest.NewLogger(t)
	logger := zapr.NewLogger(zapLogger)

	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "testdata")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a temporary input file
	inputPath := tempDir + "/input.md"
	inputContent := []byte(`# Test Input

This is a test input file.`)
	err = os.WriteFile(inputPath, inputContent, 0o644)
	if err != nil {
		t.Fatalf("Failed to create test input file: %v", err)
	}

	// Change working directory to the temp directory
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

	// Create testdata directory
	err = os.Mkdir("testdata", 0o755)
	if err != nil {
		t.Fatalf("Failed to create testdata directory: %v", err)
	}

	// Move input file to testdata directory
	err = os.Rename(inputPath, "testdata/input.md")
	if err != nil {
		t.Fatalf("Failed to move input file: %v", err)
	}

	// Redirect stdout to capture output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err = Hello(logger, false)
	if err != nil {
		t.Fatalf("Hello() error = %v", err)
	}

	// Restore stdout
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
}