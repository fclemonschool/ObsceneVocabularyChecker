package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func makeItLouder(original string) string {
	return strings.ToUpper(original) // return LOUDER line
}

// DO NOT EDIT
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(makeItLouder(scanner.Text()))
}
