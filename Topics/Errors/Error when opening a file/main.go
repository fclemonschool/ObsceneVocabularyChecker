package main

import (
	"log"
	"os"
)

func checkFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("open nonexisting_file.txt: no such file or directory")
		return
	}
	defer file.Close()
}
