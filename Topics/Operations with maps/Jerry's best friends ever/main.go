package main

import "fmt"

func createSet() map[string]bool {
	// fix the type of the bestFriends map values.
	bestFriends := map[string]bool{
		"George Costanza": true,
		"Elaine Benes":    true,
		"Cosmo Kramer":    true,
	}
	fmt.Println()
	return bestFriends // do not delete this line!
}
