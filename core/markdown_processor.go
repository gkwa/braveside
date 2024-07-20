package core

import (
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type MarkdownProcessor struct {
	md goldmark.Markdown
}

func NewMarkdownProcessor() *MarkdownProcessor {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.DefinitionList, meta.Meta),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)
	return &MarkdownProcessor{md: md}
}

func (mp *MarkdownProcessor) parse(input []byte) (ast.Node, map[string]interface{}) {
	context := parser.NewContext()
	doc := mp.md.Parser().Parse(text.NewReader(input), parser.WithContext(context))
	metaData := meta.Get(context)
	return doc, metaData
}
