package core

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark/ast"
	"gopkg.in/yaml.v3"
)

type DefaultASTPrinter struct{}

func (p *DefaultASTPrinter) PrintAST(doc ast.Node, input []byte) {
	fmt.Println("AST structure:")
	printNode(doc, input, 0)
}

type DefaultFrontMatterProcessor struct{}

func (p *DefaultFrontMatterProcessor) ProcessFrontMatter(metaData map[string]interface{}) (string, error) {
	var frontMatterBuf bytes.Buffer
	encoder := yaml.NewEncoder(&frontMatterBuf)
	encoder.SetIndent(2)
	if err := encoder.Encode(metaData); err != nil {
		return "", err
	}
	encoder.Close()
	return frontMatterBuf.String(), nil
}

type DefaultMarkdownRenderer struct{}

func (r *DefaultMarkdownRenderer) RenderMarkdown(doc ast.Node, source []byte) ([]byte, error) {
	var contentBuf bytes.Buffer
	renderMarkdown(&contentBuf, doc, source, 0)
	return contentBuf.Bytes(), nil
}
