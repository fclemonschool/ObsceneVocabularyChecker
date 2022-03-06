package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder

	var c rune // what type allows us to read Unicode characters?
	fmt.Scanln(&c)

	for i := 0; i < 3; i++ {
		b.WriteRune(c) // what function allows us to write Unicode characters?
	}

	// This line outputs the Unicode character do not delete it!
	fmt.Println(b.String())
}
