package dataParses

import (
	"fmt"
	"strings"
)

// checkData est une fonction privée qui est utilisé seulement dans la fonction DataParses.
// Elle permet de vérifier si le format attendu des données dans le fichier .txt est le bon.
func checkData(ants int, rooms, links []string, startEndRoom map[string][]string) (err error) {

	err = findRoom(rooms, links)
	if err != nil {
		return err
	}

	err = checkAnts(ants)
	if err != nil {
		return err
	}

	err = checkRooms(rooms)
	if err != nil {
		return err
	}

	err = checkLinks(links)
	if err != nil {
		return err
	}

	err = checkStartEndRoom(startEndRoom)
	if err != nil {
		return err
	}
	return nil
}

// La fonction findRoom vérifie si les rooms exite bien.
func findRoom(rooms, links []string) error {
	var getRooms []string
	getRoomLinks := make(map[string]bool)

	for _, room := range rooms {
		cleanRoom := strings.Split(room, " ")
		getRooms = append(getRooms, cleanRoom[0])
	}

	for _, link := range links {
		cleanLink := strings.Split(link, "-")
		getRoomLinks[cleanLink[0]] = false
		getRoomLinks[cleanLink[1]] = false
	}

	if len(getRooms) < len(getRoomLinks) {
		return fmt.Errorf("Tunnel impossible: room inexistant :  %d rooms au lieu de %d\n", len(getRooms), len(getRoomLinks))
	}

	for _, getRoom := range getRooms {
		for getRoomLink := range getRoomLinks {
			if getRoom == getRoomLink {
				getRoomLinks[string(getRoomLink)] = true
			}
		}
	}

	for room, ok := range getRoomLinks {
		if ok == false {
			return fmt.Errorf("Il y a un tunnel impossible car il y a une room qui n'existe pas: rooms %s\n", room)
		}
	}

	return nil
}
