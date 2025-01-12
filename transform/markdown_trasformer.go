package transform

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strings"
)

type MarkdownTransform struct{}

// NewHTMLToMarkdown creates a new HTMLToMarkdown transformer.
func NewMarkdownTransform() *MarkdownTransform {
	return &MarkdownTransform{}
}

// Transform converts HTML to Markdown.
func (t *MarkdownTransform) Transform(data interface{}) (string, error) {
	return StructToMarkdown(data)
}

func StructToMarkdown(item interface{}) (string, error) {
	// Convert struct to map
	var data map[string]interface{}
	err := mapstructure.Decode(item, &data)
	if err != nil {
		return "", err
	}

	// Generate Markdown
	var builder strings.Builder
	for key, value := range data {
		builder.WriteString(fmt.Sprintf("**%s:**\n", key))

		switch v := value.(type) {
		case string:
			builder.WriteString(fmt.Sprintf("%s\n\n", v))
		case []string:
			for _, entry := range v {
				builder.WriteString(fmt.Sprintf("- %s\n", entry))
			}
			builder.WriteString("\n")
		}
	}

	return builder.String(), nil
}
