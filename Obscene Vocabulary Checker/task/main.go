package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var wordMap = make(map[string]string)

func main() {
	// write your code here
	openTabooWordFile()
	checkTabooWord()
}

func openTabooWordFile() {
	var fileName string
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordMap[strings.ToUpper(scanner.Text())] = scanner.Text()
	}

	if scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func checkTabooWord() {
	for true {
		var word string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			word = scanner.Text()
			if word == "exit" {
				fmt.Println("Bye!")
				break
			}
			censorWord(word)
		}
	}
}

func censorWord(word string) {
	if wordMap[strings.ToUpper(word)] == "" {
		fmt.Print(word)
	} else {
		var result string
		for i := 0; i < len(word); i++ {
			result += fmt.Sprint("*")
		}
		fmt.Printf("%v\n", result)
	}
}
