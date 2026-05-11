package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadFile lit un fichier ligne par ligne et retourne un tableau de chaînes.
func ReadFile(fileName string) ([]string, error) {
	var lines []string

	file, err := os.Open("../../tests/" + fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

// DeleteOptions supprime du tableau les lignes contenant les mots
// "##start" ou "##end".
// La fonction retourne un nouveau tableau sans ces options.
func DeleteOptions(lines []string) ([]string, error) {
	var cleanLines []string
	var options []string
	end := "##end"
	start := "##start"

	for _, l := range lines {
		if strings.Contains(l, start) || strings.Contains(l, end) {
			options = append(options, l)
		}
	}
	if len(options) < 2 {
		return nil, fmt.Errorf("Option invalide. Exemple de format : %v, %v\n", start, end)
	}

	for _, line := range lines {
		// Si la ligne ne contient pas "start" ni "end",
		// elle est conservée dans le nouveau tableau.
		if !(strings.Contains(line, "#")) {
			cleanLines = append(cleanLines, line)
		}
	}

	return cleanLines, nil
}
