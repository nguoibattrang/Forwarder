package extractor

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type ExtractConfluenceItem struct {
	Url      string   `mapstructure:"URL"`
	Title    string   `mapstructure:"Title"`
	Users    []string `mapstructure:"Users"`
	Comments []string `mapstructure:"Comments"`
	Content  string   `mapstructure:"Content"`
}

func ExtractConfluenceHTML(htmlData string) (*ExtractConfluenceItem, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlData))
	if err != nil {
		return nil, err
	}
	title := extractConfluenceTitle(doc)
	if title == "" {
		return nil, errors.New("title is empty")
	}
	return &ExtractConfluenceItem{
		Title:    title,
		Users:    extractConfluenceAuthor(doc),
		Comments: extractConfluenceComments(doc),
		Content:  extractConfluenceContent(doc),
	}, nil

}

func extractConfluenceTitle(doc *goquery.Document) string {
	return doc.Find("#summary-val").Text()
}

func extractConfluenceAuthor(doc *goquery.Document) []string {
	var users []string
	assignee := strings.TrimSpace(doc.Find("#assignee-val .user-hover").Text())
	reporter := strings.TrimSpace(doc.Find("#reporter-val .user-hover").Text())
	if assignee != "" {
		users = append(users, assignee)
	}
	if reporter != "" {
		users = append(users, reporter)
	}

	return users
}

func extractConfluenceComments(doc *goquery.Document) []string {
	var comments []string
	doc.Find(".comment").Each(func(i int, s *goquery.Selection) {
		comments = append(comments, s.Text())
	})
	return comments
}

func extractConfluenceContent(doc *goquery.Document) string {
	return doc.Find("#summary-val").Text()
}
