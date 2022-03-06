package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// What split function allows us to read character by character?
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		c := scanner.Text()
		fmt.Println(c)
	}
}
