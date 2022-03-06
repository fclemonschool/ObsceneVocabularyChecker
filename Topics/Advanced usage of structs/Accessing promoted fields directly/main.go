package main

import "fmt"

type Player struct {
	playerName     string
	playerAge      int
	playerPosition string
}

type FootballTeam struct {
	teamName  string
	coachName string
	Player    // add 'Player' as an anonymous field within the FootballTeam struct.
}

func solve() FootballTeam {
	// do not delete the 'bayern' FootballTeam declaration!
	var bayern FootballTeam
	bayern.teamName = "Bayern"
	bayern.coachName = "Nagelsmann"
	fmt.Sprint()
	// assign the values below to the promoted fields
	bayern.playerName = "Neuer"          // assign the value 'Neuer' here
	bayern.playerAge = 35                // assign the value 35 here
	bayern.playerPosition = "Goalkeeper" // assign the value 'Goalkeeper' here

	return bayern // do not delete this line!
}
