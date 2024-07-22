package core

import (
	"github.com/yuin/goldmark/ast"
)

type FrontMatterProcessor interface {
	ProcessFrontMatter(metaData map[string]interface{}) (string, error)
}

type MarkdownRenderer interface {
	RenderMarkdown(doc ast.Node, source []byte) ([]byte, error)
}
