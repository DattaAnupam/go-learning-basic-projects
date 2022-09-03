package main

import (
	"fmt"
)

func main() {
	// Way 3 of declaring Map
	// Only declaration. We declare in this way when we are going to put values
	// in future executions.
	colors := make(map[string]string)
	fmt.Println(colors)
}
