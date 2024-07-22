package core

import (
	"github.com/yuin/goldmark/ast"
)

type ASTPrinter interface {
	PrintAST(doc ast.Node, input []byte)
}

type FrontMatterProcessor interface {
	ProcessFrontMatter(metaData map[string]interface{}) (string, error)
}

type MarkdownRenderer interface {
	RenderMarkdown(doc ast.Node, source []byte) ([]byte, error)
}

type NoopASTPrinter struct{}

func (p *NoopASTPrinter) PrintAST(doc ast.Node, input []byte) {}
