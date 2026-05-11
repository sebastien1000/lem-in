package graph

import "lem-in/internal/config"

// Parcourt tous les liens pour construire le graphe
func GraphLinks(links [][]string) map[string][]string {
	graphs := map[string][]string{}

	for _, link := range links {
		graphs[link[0]] = append(graphs[link[0]], link[1])
		graphs[link[1]] = append(graphs[link[1]], link[0])
	}

	return graphs
}

/*
Retourn tout les chemins possible
*/
func GraphDFS(startend, graph map[string][]string, path []string, current string) {
	newPath := append([]string{}, path...)
	newPath = append(newPath, current)

	if current == startend["##end"][0] {
		config.FinalPaths = append(config.FinalPaths, newPath)
		return
	}

	for _, next := range graph[current] {
		if !isVisited(next, newPath) {
			GraphDFS(startend, graph, newPath, next)
		}
	}
}

// Vérifie si le chemin a été emprunté.
func isVisited(tunel string, path []string) bool {
	for _, cur := range path {
		if cur == tunel {
			return true
		}
	}

	return false
}
