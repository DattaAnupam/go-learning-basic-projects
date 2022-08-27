package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Create and return a list of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "King", "Queen"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardSuit+" of "+cardValue)
		}
	}

	return cards
}

// prints the card
// any variable of type 'deck' have access to this function
func (d deck) printCard() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
