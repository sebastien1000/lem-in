package dataParses

import (
	"errors"
	"fmt"
	"strings"
)

// checkLinks est une fonction privée qui est utilé dans la fonction checkData.
// Elle vérifie le si le format des links est bon.
func checkLinks(links []string) error {

	if len(links) == 0 {
		return errors.New("Links est vide\n")
	}

	for i := 0; i < len(links); i++ {
		parts1 := strings.Split(links[i], "-")
		if len(parts1) != 2 {
			return fmt.Errorf("Seules deux rooms doivent définir un tunnel (link) : %s\n", links[i])
		}

		for j := i + 1; j < len(links); j++ {
			parts2 := strings.Split(links[j], "-")
			if len(parts2) != 2 {
				return fmt.Errorf("Seules deux rooms doivent définir un tunnel (link): %s\n", links[j])
			}

			if links[i] == links[j] {
				return fmt.Errorf("Doublon de tunnel (link) : %s. Index : %d et %d\n", links[i], i, j)
			}

			if parts1[0] == parts2[1] && parts1[1] == parts2[0] {
				return fmt.Errorf("Doublon inversé de tunnel (link) : %s et %s. Index : %d et %d\n", links[i], links[j], i, j)
			}
		}
	}
	return nil
}
