package parses

import (
	"strings"
)

// ParseLinks() récupère les rooms liées ensemble
func ParseLinks(lines []string) [][]string {
	var links [][]string

	for _, line := range lines {
		links = append(links, strings.Split(line, "-"))
	}
	return links
}
