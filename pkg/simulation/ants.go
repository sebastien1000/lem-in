package simulation

import (
	"fmt"
	"lem-in/internal/models"
	"math"
	"strings"
)

// MoveAllAnts() déplace les fourmis une par une et les répartit dans les différents chemins.
func MoveAllAnts(totalAnts int, finalPaths []models.PathInfo) {

	// --- 1. Distribution des fourmis (inchangée, elle est correcte) ---
	for i := 0; i < totalAnts; i++ {
		bestIndex := 0
		minTime := math.MaxInt32

		for j, p := range finalPaths {
			finishTime := (len(p.Path) - 1) + p.AntNbr
			if finishTime < minTime {
				minTime = finishTime
				bestIndex = j
			}
		}
		finalPaths[bestIndex].AntNbr++
	}

	// --- 2. Création des fourmis en alternant les chemins ---
	// Au lieu de créer toutes les fourmis du chemin 0 puis du chemin 1,
	// on alterne : fourmi 1 sur chemin 0, fourmi 2 sur chemin 1, etc.
	// Ainsi L1 et L2 partent ensemble dès le tour 1.
	var ants []*models.Ant
	id := 1

	// Trouver le max d'AntNbr pour savoir combien de "rounds" d'affectation
	maxAnts := 0
	for _, p := range finalPaths {
		if p.AntNbr > maxAnts {
			maxAnts = p.AntNbr
		}
	}

	for round := 0; round < maxAnts; round++ {
		for _, p := range finalPaths {
			if round < p.AntNbr {
				ants = append(ants, &models.Ant{
					ID:    id,
					Path:  p.Path,
					Index: 0,
				})
				id++
			}
		}
	}

	// --- 3. Simulation du déplacement ---
	// On marque aussi la salle START comme occupée pour éviter
	// que deux fourmis soient au même endroit
	finished := 0

	for finished < totalAnts {
		var moves []string
		// occupied représente les salles qui seront occupées EN FIN de tour
		occupied := make(map[string]bool)

		for _, ant := range ants {
			if ant.Index >= len(ant.Path)-1 {
				continue
			}

			nextIndex := ant.Index + 1
			nextRoom := ant.Path[nextIndex]
			isEnd := nextIndex == len(ant.Path)-1

			occupyKey := nextRoom
			if isEnd {
				occupyKey = "end:" + ant.Path[nextIndex-1]
			}

			if !occupied[occupyKey] {
				ant.Index = nextIndex
				moves = append(moves, fmt.Sprintf("L%d-%s", ant.ID, nextRoom))
				occupied[occupyKey] = true
				if isEnd {
					finished++
				}
			}
		}

		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}
