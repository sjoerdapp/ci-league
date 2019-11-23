package league

import "regexp"

var (
	coAuthorRegex = regexp.MustCompile(`Co-authored-by:.*<(.*)>`)
)

func extractCoAuthor(message string) string {
	matches := coAuthorRegex.FindStringSubmatch(message)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
