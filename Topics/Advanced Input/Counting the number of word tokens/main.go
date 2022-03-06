package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Set the correct split function below:
	scanner.Split(bufio.ScanWords)

	// Write the for loop below with the correct scanner statement
	count := 0
	for scanner.Scan() {
		count++
	}
	fmt.Println(count)
}
