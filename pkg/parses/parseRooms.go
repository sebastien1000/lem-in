package parses

import (
	"lem-in/internal/models"
	"strconv"
	"strings"
)

// ParseRooms() sépare la room des ses coordonnées
func ParseRooms(lines []string) []*models.Room {
	var rooms []*models.Room

	for i := range lines {

		// Découpe la ligne en trois éléments : nom, x, y
		numbers := strings.Split(lines[i], " ")
		x, _ := strconv.Atoi(numbers[1])
		y, _ := strconv.Atoi(numbers[2])

		// Création de la room et ajout dans le tableau
		rooms = append(rooms, &models.Room{Name: numbers[0], X: x, Y: y})
	}
	return rooms
}
