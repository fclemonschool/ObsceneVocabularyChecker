package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("galley_grub.txt") // open the file here with the ioutil.ReadFile() function
	if err != nil {
		fmt.Println(err) // print the error in case there is any
	}
	// print the contents of the file here as a string
	fmt.Println(string(file))
}
