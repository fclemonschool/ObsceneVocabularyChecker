package main

import (
	"fmt"
	"math"
)

func main() {
	// Do not change the type of the variables from 'float64' to 'int'
	var totalQuestions float64
	fmt.Scanln(&totalQuestions)

	var correctAnswers float64
	fmt.Scanln(&correctAnswers)

	// Please write the grading formula within the 'score' variable below:
	score := correctAnswers / totalQuestions * 10

	// What function from the math package allows us
	// to either round up or down based on the decimal numbers?
	score = math.Round(score)

	fmt.Println(score)
}
