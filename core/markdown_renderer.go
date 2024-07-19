package core

import (
	"io"

	"github.com/yuin/goldmark/ast"
	east "github.com/yuin/goldmark/extension/ast"
)

func renderMarkdown(w io.Writer, n ast.Node, source []byte, level int) {
	switch v := n.(type) {
	case *ast.AutoLink:
		renderAutoLink(w, v, source)
	case *ast.Blockquote:
		renderBlockquote(w, v, source, level)
	case *ast.CodeSpan:
		renderCodeSpan(w, v, source, level)
	case *east.DefinitionDescription:
		renderDefinitionDescription(w, v, source, level)
	case *east.DefinitionList:
		renderDefinitionList(w, v, source, level)
	case *east.DefinitionTerm:
		renderDefinitionTerm(w, v, source, level)
	case *ast.Document:
		renderDocument(w, v, source, level)
	case *ast.Emphasis:
		renderEmphasis(w, v, source, level)
	case *ast.FencedCodeBlock:
		renderFencedCodeBlock(w, v, source)
	case *ast.Heading:
		renderHeading(w, v, source, level)
	case *ast.Image:
		renderImage(w, v, source)
	case *ast.Link:
		renderLink(w, v, source, level)
	case *ast.List:
		renderList(w, v, source, level)
	case *ast.ListItem:
		renderListItem(w, v, source, level)
	case *ast.Paragraph:
		renderParagraph(w, v, source, level)
	case *ast.String:
		renderString(w, v)
	case *east.Strikethrough:
		renderStrikethrough(w, v, source, level)
	case *east.Table:
		renderTable(w, v, source)
	case *east.TaskCheckBox:
		renderTaskCheckBox(w, v)
	case *ast.Text:
		renderText(w, v, source)
	case *ast.TextBlock:
		renderTextBlock(w, v, source, level)
	case *ast.ThematicBreak:
		renderThematicBreak(w)
	default:
		renderDefault(w, n, source)
	}
}
