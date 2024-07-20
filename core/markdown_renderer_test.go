package core

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func TestRenderMarkdown(t *testing.T) {
	input := []byte(`---
title: Test Document
author: John Doe
---

# Heading 1

## Heading 2

This is a paragraph with **bold** and *italic* text.

- List item 1
- List item 2
  - Nested item 1
  - Nested item 2

1. Ordered item 1
2. Ordered item 2

> This is a blockquote.

[Link](https://example.com)

![Image](https://example.com/image.jpg)

` + "```go" + `
package main

func main() {
    println("Hello, World!")
}
` + "```" + `

| Column 1 | Column 2 |
|----------|----------|
| Cell 1   | Cell 2   |
| Cell 3   | Cell 4   |
`)

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.DefinitionList, meta.Meta),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

	context := parser.NewContext()
	doc := md.Parser().Parse(text.NewReader(input), parser.WithContext(context))

	var buf bytes.Buffer
	renderMarkdown(&buf, doc, input, 0)

	expected := `# Heading 1

## Heading 2

This is a paragraph with **bold** and *italic* text.

- List item 1
- List item 2
  - Nested item 1
  - Nested item 2

1. Ordered item 1
2. Ordered item 2

> This is a blockquote.

[Link](https://example.com)

![Image](https://example.com/image.jpg)

` + "```go" + `
package main

func main() {
    println("Hello, World!")
}
` + "```" + `

| Column 1 | Column 2 |
|----------|----------|
| Cell 1   | Cell 2   |
| Cell 3   | Cell 4   |
`

	assert.Equal(t, expected, buf.String())
}
