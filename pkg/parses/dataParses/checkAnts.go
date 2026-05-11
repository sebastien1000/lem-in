package dataParses

import "fmt"

// checkAnts est une fonction privée qui est utilé dans la fonction checkData.
// Elle vérifie le si le format des ants est bon
func checkAnts(ants int) error {
	if ants <= 0 {
		return fmt.Errorf("Nombre de fourmis invalide: %v. Format attendu: [fournis > 0]\n", ants)
	}
	return nil
}
