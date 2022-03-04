package main

import "fmt"

func main() {
	// do not change the contents of the nominees map!
	nominees := map[string]string{
		"Lionel Messi":       "Paris Sant-Germain",
		"Robert Lewandowski": "Bayern Munich",
		"Jorginho":           "Chelsea",
		"Cristiano Ronaldo":  "Manchester United",
		"Karim Benzema":      "Real Madrid",
	}

	fmt.Println(len(nominees)) // print the length of the nominees map here
}
