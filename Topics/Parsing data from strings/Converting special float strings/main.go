package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// Do not change or delete these two lines below!
	var stringNum string
	fmt.Scanln(&stringNum)

	// Write in the '?' below the missing code to complete the program
	value, err := strconv.ParseFloat(stringNum, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Please add the missing code to print the converted type and its value below:
	fmt.Printf("%T %v", value, value)
}
