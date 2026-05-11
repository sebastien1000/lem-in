package dataParses

import "fmt"

// Vérifie si les entrées et les sorties existent.
func checkStartEndRoom(startEndRoom map[string][]string) error {
	clesAVerifier := []string{"##start", "##end"}

	for _, cle := range clesAVerifier {
		_, existe := startEndRoom[cle]
		if existe == true {
			continue
		} else {
			return fmt.Errorf("manque : %v\n", cle)
		}

	}
	return nil
}
