package dataParses

import (
	"fmt"
	"lem-in/pkg/utils"
	"strconv"
	"strings"
)

/*
	DataParses retourne :

(1) le nombre de fourmis (ants),
(2) une slice de rooms,
(3) une slice de links,
(4) une slice de rooms de départ et d'arrivée,
(5) ainsi qu’une erreur si le format attendu n'est pas le bon.
*/
func DataParses(lines []string) (ants int, rooms, links []string, startEndRoom map[string][]string, err error) {
	var newlines []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		newlines = append(newlines, line)
	}

	startEndRoom = getOptions(newlines)

	newlines, err = utils.DeleteOptions(newlines)
	if err != nil {
		return 0, nil, nil, nil, err
	}

	ants, err = strconv.Atoi(newlines[0])
	if err != nil {
		return 0, nil, nil, nil, fmt.Errorf("Nombre de fourmis invalide: %v\n", err)
	}

	for i := 1; i < len(newlines); i++ {
		newlines[i] = strings.TrimSpace(newlines[i])

		if newlines[i] == "" {
			continue
		}

		if strings.Contains(newlines[i], " ") && !strings.Contains(newlines[i], "-") {
			rooms = append(rooms, newlines[i])
		} else if strings.Contains(newlines[i], "-") && !strings.Contains(newlines[i], " ") {
			links = append(links, newlines[i])
		} else {
			return 0, nil, nil, nil, fmt.Errorf("La liaison des tunnels (links) \"%s\" doit être représentée avec ce caractère \"-\" au lieu de \"%c\"\n", newlines[i], newlines[i][1])
		}
	}

	err = checkData(ants, rooms, links, startEndRoom)
	if err != nil {
		return 0, nil, nil, nil, err
	}

	return ants, rooms, links, startEndRoom, nil
}

// getOptions() est une fonction privée qui se trouve dans DataParses et récupère la rooms ##start et ##end.
func getOptions(lines []string) (startEndRoom map[string][]string) {
	startEndRoom = make(map[string][]string)
	for i, line := range lines {
		switch {
		case line == "##start":
			if i+1 < len(lines) {
				startEndRoom[line] = append(startEndRoom[line], splitFirstField(lines[i+1]))

			}
		case line == "##end":
			if i+1 < len(lines) {
				startEndRoom[line] = append(startEndRoom[line], splitFirstField(lines[i+1]))
			}
		}
	}

	return startEndRoom
}

// splitFirstField retourne le premier mot d'une ligne (avant le premier espace).
func splitFirstField(line string) string {
	for i, c := range line {
		if c == ' ' {
			return line[:i]
		}
	}
	return line
}
