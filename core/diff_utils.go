package core

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/go-logr/logr"
)

func CompareDiff(logger logr.Logger, file1, file2 io.Reader) (string, error) {
	tempFile1, err := createTempFile(file1)
	if err != nil {
		return "", fmt.Errorf("failed to create temp file for file1: %w", err)
	}
	defer os.Remove(tempFile1.Name())

	tempFile2, err := createTempFile(file2)
	if err != nil {
		return "", fmt.Errorf("failed to create temp file for file2: %w", err)
	}
	defer os.Remove(tempFile2.Name())

	opts := []string{
		"--unified",
		"--ignore-blank-lines",
		"--ignore-all-space",
	}
	args := append(opts, tempFile1.Name(), tempFile2.Name())
	cmd := exec.Command("diff", args...)
	diff, err := cmd.CombinedOutput()
	if err != nil && err.(*exec.ExitError).ExitCode() != 1 {
		return "", fmt.Errorf("diff command failed: %w", err)
	}

	var result string
	if len(diff) > 0 {
		result = fmt.Sprintf("Differences found:\ndiff%s", string(diff))
	} else {
		result = "No differences found between input and output"
		logger.Info(result)
	}

	return result, nil
}

func createTempFile(content io.Reader) (*os.File, error) {
	tempFile, err := os.CreateTemp("", "diff-*")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(tempFile, content)
	if err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, err
	}

	tempFile.Close()
	return tempFile, nil
}
