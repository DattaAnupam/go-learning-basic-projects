package main

import (
	"fmt"
)

func main() {
	colors := make(map[string]string)

	colors["firstColor"] = "Green"
	colors["secondColor"] = "Yellow"
	colors["thirdColor"] = "Red"

	fmt.Println("Colors before deleting first color")
	fmt.Println(colors)

	// delete value from map
	// call delete function, pass the map name and key
	delete(colors, "firstColor")

	fmt.Println("Colors after deleting first color")
	fmt.Println(colors)
}
