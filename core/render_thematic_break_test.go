package core

import (
	"bytes"
	"strings"
	"testing"
)

func TestRenderThematicBreak(t *testing.T) {
	var buf bytes.Buffer
	renderThematicBreak(&buf)
	expected := "---\n\n"
	if strings.TrimSpace(buf.String()) != strings.TrimSpace(expected) {
		t.Errorf("Expected %q, got %q", strings.TrimSpace(expected), strings.TrimSpace(buf.String()))
	}
}
