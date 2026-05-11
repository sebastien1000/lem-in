package main

import (
	"fmt"
	"lem-in/internal/config"
	"lem-in/pkg/graph"
	"lem-in/pkg/parses"
	dataParses "lem-in/pkg/parses/dataParses"
	"lem-in/pkg/simulation"
	"lem-in/pkg/utils"
	"lem-in/pkg/way"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Erreur, il doit y avoir un argument. Exemple : run main.go file.txt")
		return
	}

	filePath := os.Args[1]
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		fmt.Println("ERREUR: format de données invalide.")
		return
	}

	ants, rooms, links, startEndRoom, err := dataParses.DataParses(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 1. Affichage du fichier d'origine (requis par le sujet)
	fmt.Println(strings.Join(lines, "\n"))
	fmt.Println()

	// 2. Parsing et préparation du graphe
	roomData := parses.ParseRooms(rooms)
	var roomNames []string
	for _, r := range roomData {
		roomNames = append(roomNames, r.Name)
	}

	linkData := parses.ParseLinks(links)
	mapGraph := graph.GraphLinks(linkData)

	// 3. Lancement du DFS pour trouver le chemin
	start := startEndRoom["##start"][0]
	graph.GraphDFS(startEndRoom, mapGraph, nil, start)
	if len(config.FinalPaths) == 0 {
		fmt.Println("ERREUR: aucun chemin trouvé")
		return
	}

	finalPathsTest := way.ParsesPath(startEndRoom)

	nbrexit := utils.HowManyEnd(links, start)

	selectedpath := way.SelectPath(finalPathsTest, nbrexit)
	if len(selectedpath) == 0 {
		fmt.Println("ERREUR: aucune combinaison de chemin trouvé.")
		return
	}

	// 4. Déplacement des fourmis
	pathInfos := way.ConvertToPathInfo(selectedpath)
	simulation.MoveAllAnts(ants, pathInfos)
}
