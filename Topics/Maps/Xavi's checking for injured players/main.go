package main

import "fmt"

func main() {
	// do not change the values inside the map!
	injuredPlayers := map[string]int{
		"Aguero":  93,
		"Dembele": 45,
		"Pedri":   4,
	}

	var player string

	fmt.Scanln(&player) // take input of the player name here within the 'player' variable

	// implement the val, ok check below to verify if the player is on the map
	if val, ok := injuredPlayers[player]; ok {
		fmt.Println("The player", player, "will recover in", val, "days")
		return
	}
	fmt.Println("The player", player, "is not injured")
}
