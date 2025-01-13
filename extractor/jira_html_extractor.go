package extractor

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type ExtractJiraItem struct {
	Url          string   `mapstructure:"URL"`
	Title        string   `mapstructure:"Title"`
	Description  string   `mapstructure:"Description"`
	Users        []string `mapstructure:"Users"`
	Comments     []string `mapstructure:"Comments"`
	SubtaskLinks []string `mapstructure:"Subtask Links"`
}

func ExtractJiraHTML(htmlData string) (string, *ExtractJiraItem, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlData))
	if err != nil {
		return "", nil, err
	}
	title := extractJiraTitle(doc)
	if title == "" {
		return "", nil, errors.New("title is empty")
	}
	return title, &ExtractJiraItem{
		Title:        title,
		Description:  extractJiraDescription(doc),
		Comments:     extractJiraComments(doc),
		SubtaskLinks: extractJiraSubtask(doc),
		Users:        extractJiraAssignee(doc),
	}, nil

}

func extractJiraTitle(doc *goquery.Document) string {
	return doc.Find("#summary-val").Text()
}

func extractJiraDescription(doc *goquery.Document) string {
	return strings.TrimSpace(doc.Find("#description-val").Text())
}

func extractJiraAssignee(doc *goquery.Document) []string {
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

func extractJiraComments(doc *goquery.Document) []string {
	var comments []string
	doc.Find(".comment").Each(func(i int, s *goquery.Selection) {
		comments = append(comments, s.Text())
	})
	return comments
}

func extractJiraSubtask(doc *goquery.Document) []string {
	var tasks []string

	// Find all titles in the table
	doc.Find("#issuetable .stsummary a").Each(func(index int, item *goquery.Selection) {
		// Extract text content
		title := strings.TrimSpace(item.Text())
		tasks = append(tasks, title)
	})
	return tasks
}
