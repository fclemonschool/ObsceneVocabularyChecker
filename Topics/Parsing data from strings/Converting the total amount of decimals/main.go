package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Do not change or delete these two lines below!
	var mathConstant float64
	fmt.Scanln(&mathConstant)

	// Write in the '?' the missing arguments required to properly convert
	// all of the decimals of the 'mathConstant' variable to a string!
	val := strconv.FormatFloat(mathConstant, 'f', -1, 64)

	// This line prints the converted value, do not delete it!
	fmt.Printf("%s", val)
}
