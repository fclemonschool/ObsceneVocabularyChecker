package main

import "fmt"

func main() {
	// delete or add new defer statements below!
	var name string
	fmt.Scanln(&name)

	fmt.Printf("Your name is: %s\n", string(name))
	fmt.Printf("Reversed name: ")

	for _, char := range []rune(name) {
		defer fmt.Printf("%c", char)
	}
}
