package main

import (
    "fmt"
)

func main() {
    // help keanu fix his map declaration - do not change the values inside the map!
    keanuRoles := make(map[string]string)
    keanuRoles["Constantine"] = "John Constantine 👨🏻‍💼"
    keanuRoles["John Wick"] = "Jardani Jovanovich 🤵🏻‍♂️"
    keanuRoles["The Matrix"] = "Neo 🕵🏻‍♂️"
    keanuRoles["Speed"] = "Jack Traven 👮🏻"

    fmt.Println(keanuRoles)
}