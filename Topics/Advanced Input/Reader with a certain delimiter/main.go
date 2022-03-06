package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// This line reads the custom delimiter; do not delete it!
	var delimiter byte
	fmt.Scanf("%c", &delimiter)

	// What Reader functions take a 'delimiter' as an argument?
	s, err := reader.ReadBytes(delimiter)
	if err != nil {
		log.Println(err)
	}
	// Output the string below
	fmt.Println(string(s))
}
