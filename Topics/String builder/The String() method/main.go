package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var b strings.Builder

	// The 'scanner' lines below allow us to read a string separated by whitespaces do not delete them!
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str1 := scanner.Text()
	scanner.Scan()
	str2 := scanner.Text()

	b.WriteString(str1)
	b.WriteString(str2)

	// What is the missing method call required to print the concatenated string?
	fmt.Println(b.String())
}
