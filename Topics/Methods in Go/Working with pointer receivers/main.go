package main

import "fmt"

type Gesture struct {
	Name, Emoji string
}

// Update method with a pointer receiver to update 'Gesture' struct fields
func (g *Gesture) Update(name, emoji string) {
	g.Name = name
	g.Emoji = emoji
}

func main() {
	// Do not change the contents of the 'g' struct!
	g := Gesture{
		Name:  "OK-Hand",
		Emoji: "👌",
	}

	fmt.Printf("The current gesture is: %s %s\n", g.Name, g.Emoji)

	// Call the Update() method on the 'g' struct and pass "Thumbs-Up" and "👍" as arguments below:
	// you may copy and paste the "👍" emoji so its easier!
	g.Update("Thumbs-Up", "👍")

	fmt.Printf("The updated gesture is: %s %s\n", g.Name, g.Emoji)
}
