package main

import (
	"fmt"
)

func main() {
	// the galleyGrub map contains the Krusty Krab menu
	galleyGrub := map[string]float64{
		"Krabby Pattie 🍔":  2.49,
		"Krusty Combo 🍔🥤🍟": 3.99,
		"Coral Bits 🍟":     1.95,
	}

	galleyGrub["Krusty Dog 🌭"] = 3.99 // add the Krusty Dog 🌭 at a cost of 3.99 here.

	fmt.Println(galleyGrub)
}
