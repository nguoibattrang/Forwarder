package extractor

import "fmt"

func ExtractHTML(site string, content string) (interface{}, error) {
	switch site {
	case "jira":
		return ExtractJiraHTML(content)
	default:
		return nil, fmt.Errorf("unsupported site %s", site)
	}
}
