package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder

	// call the Write() function two times with the slice of bytes containing the phrases:
	b.Write([]byte("Aye, carumba!"))

	b.Write([]byte("Eat my shorts!"))

	// print the built string below
	fmt.Println(b.String())
}
