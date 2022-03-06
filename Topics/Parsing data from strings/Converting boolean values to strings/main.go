package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Do not change or delete these two lines below!
	var b bool
	fmt.Scanln(&b)

	// What is the correct function to convert bool types to strings?
	val := strconv.FormatBool(b)

	// This line prints the converted type and its value, do not delete it!
	fmt.Printf("%T %s", val, val)
}
