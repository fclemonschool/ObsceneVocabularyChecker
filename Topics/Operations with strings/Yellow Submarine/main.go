package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func numberOfWords(line string) int {
	return len(strings.Split(line, " ")) // number of words
}

// DO NOT EDIT
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(numberOfWords(scanner.Text()))
}
