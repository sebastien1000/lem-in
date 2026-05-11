package dataParses

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// checkRooms est une fonction privée qui est utilé dans la fonction checkData.
// Elle vérifie le si le format des rooms est bon.
func checkRooms(rooms []string) (err error) {
	if len(rooms) == 0 {
		return fmt.Errorf("Aucune room trouvée: %v\n", rooms)
	}
	for i := range rooms {
		numbers := strings.Split(rooms[i], " ")

		if len(numbers) != 3 {
			return fmt.Errorf("Room invalide: [%v]. Format attendu: [<Nom> <x> <y>]\n", rooms[i])
		}

		for k := 1; k < len(numbers); k++ {
			_, err = strconv.Atoi(numbers[k])
			if err != nil {
				return fmt.Errorf("Room invalide: %v. Format attendu: un chiffre ou un nombre au lieu de \"%v\"\n", numbers, numbers[k])
			}
		}

		for j := i + 1; j < len(rooms); j++ {
			if rooms[i] == rooms[j] {
				return fmt.Errorf("Room invalide: %v. Identiques\n", rooms[j])
			}
			roomi := strings.Split(rooms[i], " ")
			roomj := strings.Split(rooms[j], " ")
			if roomi[0] == roomj[0] {
				return fmt.Errorf("Room invalide: %c. Room identiques\n", rooms[j][0])
			}
			if slices.Equal(roomi[1:], roomj[1:]) {
				return fmt.Errorf("Room invalide: %v. Coordonnées identiques\n", rooms[j][1:])
			}
		}
	}

	return nil
}
