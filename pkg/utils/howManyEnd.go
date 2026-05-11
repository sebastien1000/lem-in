package utils

import "strings"

/*
permet de compter le nombre de chemins menant a la sortie
*/
func HowManyEnd(links []string, start string) int {
	count := 0
	for _, link := range links {
		parts := strings.Split(link, "-")
		if parts[0] == start || parts[1] == start {
			count++
		}
	}
	return count
}
