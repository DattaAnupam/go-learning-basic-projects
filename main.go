package main

import (
	"fmt"
)

func main() {
	colors := make(map[string]string)
	// put value into map
	// using key-value
	colors["firstColor"] = "Green"
	colors["secondColor"] = "Yellow"
	colors["thirdColor"] = "Red"

	fmt.Println(colors)
}
