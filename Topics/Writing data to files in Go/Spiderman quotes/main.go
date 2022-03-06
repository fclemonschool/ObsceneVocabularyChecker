package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	quote := readQuote() // today's quote
	file, err := os.Create("quote.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// TODO: write the quote to a file
	_, err = fmt.Fprint(file, quote)
	if err != nil {
		log.Fatal(err)
	}
}

func readQuote() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	_ = scanner.Scan()
	return scanner.Text()
}
