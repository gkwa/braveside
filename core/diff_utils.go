package core

import (
	"fmt"
	"os/exec"

	"github.com/go-logr/logr"
)

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
