package main

import "fmt"

// do not change the name or the contents of the 'Band' struct!
type Band struct {
	Name    string
	Members []string
	Genre   string
}

func main() {
	// do not change the contents or the name of the 'greenDay' struct!
	greenDay := Band{
		Name:    "Green Day",
		Members: []string{"Billie Joe Armstrong", "Mike Dirnt", "Tré Cool"},
		Genre:   "Punk rock",
	}

	// print each one of the individual fields of the 'greenDay' struct here
	fmt.Println(greenDay.Name)
	fmt.Println(greenDay.Members)
	fmt.Println(greenDay.Genre)
}
