package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// prints the card
// any variable of type 'deck' have access to this function
func (d deck) printCard() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
