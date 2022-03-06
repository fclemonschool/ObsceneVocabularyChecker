package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Do not delete the splitFloat() function!
func splitFloat(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanWords(data, atEOF)
	if err == nil && token != nil {
		_, err = strconv.ParseFloat(string(token), 64)
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Set 'splitFloat' as the split function below
	scanner.Split(splitFloat)

	// Write the for loop to scan and then output the scanned data
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
