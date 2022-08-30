package main

import (
	"fmt"
	"os"
	"strings"
)

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

// deal() returns two 'deck' type values, which are of type []string
// d is the main deck on which it will operate.
// handSize is from which it will divide the deck.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// deckToString() returns string representation of the deck
func (d deck) deckToString() string {
	// convert deck to slice of string
	// join each element inside the slice of string, seperate them with a comma (,)
	return strings.Join([]string(d), ",")
}

// Save deck of card into the hard drive, with 'filename' and default permissions
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.deckToString()), 0666) // 0666: anyone can read or write to the file
}
