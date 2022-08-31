package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
	age       int
}

func main() {
	anupam := Person{
		firstname: "Anupam",
		lastname:  "Datta",
		age:       28,
	}

	fmt.Println(anupam)
}
