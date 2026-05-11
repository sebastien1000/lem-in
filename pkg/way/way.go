package way

import (
	"lem-in/internal/config"
	"lem-in/internal/models"
	"lem-in/pkg/utils"
	"sort"
)

/*
prendre un nombre de chemin a trouver (le nbr de tunel qui menne a la sortie)
prend tout les chemin possible (fptest) du plus petit au plus grand
ressort les chemins voulu
*/
func SelectPath(fptest [][]string, nbr int) [][]string {
	var bestSet [][]string

	for i := 0; i < len(fptest); i++ {
		currentSet := [][]string{fptest[i]}

		for j := 0; j < len(fptest); j++ {
			if i == j {
				continue
			}
			if !isCross(currentSet, fptest[j]) {
				currentSet = append(currentSet, fptest[j])
			}
			if len(currentSet) == nbr {
				return currentSet
			}
		}

		// Garde le meilleur ensemble trouvé jusqu'ici
		if len(currentSet) > len(bestSet) {
			bestSet = currentSet
		}
	}

	return bestSet // jamais nil, au pire le premier chemin seul
}

/*
retourne false si aucune salle n'as de room qui entre en collision
*/
func isCross(selectedpath [][]string, path []string) bool {
	// salles internes de path
	inner2 := path[1 : len(path)-1]

	// set pour path
	set2 := make(map[string]bool)
	for _, room := range inner2 {
		set2[room] = true
	}

	// comparer path avec chaque path de selectedpath
	for _, p := range selectedpath {
		if len(p) < 2 {
			continue
		}

		inner1 := p[1 : len(p)-1]

		for _, room := range inner1 {
			if set2[room] {
				return true // salle commune trouvée
			}
		}
	}

	return false // aucune salle commune
}

/*
trier les chemin en regardant ceux qui on une end
- du plus petit au plus grand
*/
func ParsesPath(startEndRoom map[string][]string) [][]string {
	var finalPaths [][]string

	for _, finalPath := range config.FinalPaths {
		if utils.Contains(finalPath, startEndRoom["##start"][0]) &&
			utils.Contains(finalPath, startEndRoom["##end"][0]) {
			finalPaths = append(finalPaths, finalPath)
		}
	}

	// Tri du plus petit chemin au plus grand
	sort.Slice(finalPaths, func(i, j int) bool {
		return len(finalPaths[i]) < len(finalPaths[j])
	})

	return finalPaths
}

/*
s
prend path et le convertie en pathInfo
*/
func ConvertToPathInfo(paths [][]string) []models.PathInfo {
	result := make([]models.PathInfo, len(paths))
	for i, p := range paths {
		result[i] = models.PathInfo{
			Path:   p,
			AntNbr: 0,
		}
	}
	return result
}
