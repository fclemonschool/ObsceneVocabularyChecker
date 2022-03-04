package main

import (
	"fmt"
	"os"
)

func main() {
	// Write the code below to open a file named "songs.txt"
	file, err := os.Open("songs.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
}
