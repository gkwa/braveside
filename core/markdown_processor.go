package core

import (
	"fmt"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
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

type MarkdownProcessor struct {
	md                   goldmark.Markdown
	astPrinter           ASTPrinter
	frontMatterProcessor FrontMatterProcessor
	markdownRenderer     MarkdownRenderer
}

func NewMarkdownProcessor(astPrinter ASTPrinter, frontMatterProcessor FrontMatterProcessor, markdownRenderer MarkdownRenderer) *MarkdownProcessor {
	return &MarkdownProcessor{
		md: goldmark.New(
			goldmark.WithExtensions(extension.GFM, extension.DefinitionList, meta.Meta),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
		),
		astPrinter:           astPrinter,
		frontMatterProcessor: frontMatterProcessor,
		markdownRenderer:     markdownRenderer,
	}
}

func (mp *MarkdownProcessor) ProcessMarkdown(input []byte) ([]byte, error) {
	context := parser.NewContext()
	doc := mp.md.Parser().Parse(text.NewReader(input), parser.WithContext(context))
	metaData := meta.Get(context)

	if mp.astPrinter != nil {
		mp.astPrinter.PrintAST(doc, input)
	}

	renderedContent, err := mp.markdownRenderer.RenderMarkdown(doc, input)
	if err != nil {
		return nil, err
	}

	if len(metaData) == 0 {
		return renderedContent, nil
	}

	processedFrontMatter, err := mp.frontMatterProcessor.ProcessFrontMatter(metaData)
	if err != nil {
		return nil, err
	}

	output := []byte(fmt.Sprintf("---\n%s---\n\n%s", processedFrontMatter, string(renderedContent)))
	return output, nil
}
