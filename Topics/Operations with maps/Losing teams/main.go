package main

import "fmt"

func main() {
	groupH := map[string]int{
		"Chelsea":  12,
		"Juventus": 12,
		"Zenit":    4,
		"Malmo":    1,
	}

	// delete the losing teams from the 'groupH' map here
	delete(groupH, "Zenit")
	delete(groupH, "Malmo")

	// do not change the text in the Println statement below!
	fmt.Println("The two teams that will pass to the knockout stage are:")

	// implement the for...range loop here to print the name of the qualified teams
	for key, _ := range groupH {
		fmt.Println(key)
	}
}
