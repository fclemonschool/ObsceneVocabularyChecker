package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Scanln(&num)

	var n float64
	fmt.Scanln(&n)

	// calculate the 'n' root of 'num' using the math.Pow() function below
	root := math.Pow(num, 1/n)

	fmt.Printf("%.2f", root)
}
