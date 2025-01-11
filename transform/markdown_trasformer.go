package transform

import (
	"github.com/gomarkdown/markdown"
)

type HTMLToMarkdown struct{}

// NewHTMLToMarkdown creates a new HTMLToMarkdown transformer.
func NewHTMLToMarkdown() *HTMLToMarkdown {
	return &HTMLToMarkdown{}
}

// Transform converts HTML to Markdown.
func (t *HTMLToMarkdown) Transform(data string) string {
	md := markdown.ToHTML([]byte(data), nil, nil)
	return string(md)
}
