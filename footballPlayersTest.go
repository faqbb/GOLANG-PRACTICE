package main

import (
	"fmt"
	"strings"
)

type Player struct {
	Name       string
	Identifier int
}

func main() {
	// Lista de jugadores actual
	players := []string{
		"John Smith", "David Johnson", "Michael Garcia", "Sarah Williams", "Daniel Martinez", "Emily Brown", " James Rodriguez", "Sophia Lee", "Lucas Oliveira", "Olivia Taylor", "Mateo Hernandez", "Emma Wilson", "Gabriel Silva"}
	newPlayers := []string{"NewPlayerYeah", "NewPlayer2", "NewPlayer3"}

	// elimina los primeros 3
	players = players[3:]

	// agregar
	for _, name := range newPlayers {
		players = append(players, name)
	}

	// mostrar lista
	fmt.Println("Lista completa: ")
	for i, player := range players {
		fmt.Printf("%d. %s\n", i+1001, sanitizeName(player))
	}
}

// eliminar nombre
func sanitizeName(name string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == ' ' {
			return r
		}
		return -1
	}, name)
}
