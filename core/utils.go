package core

import (
	"strings"
)

func indent(level int) string {
	return strings.Repeat("  ", level)
}
