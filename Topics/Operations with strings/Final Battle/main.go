package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func finalWords(obiwanWords, anakinWords string) string {
	// How to make line ENRAGED and calm?
	// Obi-Wan says first, then goes Anakin
	return strings.ToLower(obiwanWords) + " " + strings.ToUpper(anakinWords)
}

// DO NOT EDIT
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	obiwanWords := scanner.Text()
	scanner.Scan()
	anakinWords := scanner.Text()

	fmt.Println(finalWords(obiwanWords, anakinWords))
}
