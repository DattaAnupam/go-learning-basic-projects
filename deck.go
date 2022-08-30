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
	// convert deck to []string
	ss := []string(d)
	// convert []string to string
	s := strings.Join(ss, ",")
	// s := strings.Join([]string(d), ",")
	return s
}

// Save deck of card into the hard drive, with 'filename' and default permissions
func (d deck) saveToFile(filename string) error {
	// convert the deck into string
	s := d.deckToString()
	// convert the string to []byte
	sb := []byte(s)
	// bs := []byte(d.deckToString())
	return os.WriteFile(filename, sb, 0666) // 0666: anyone can read or write to the file
}

func getFromFile(filename string) deck {
	// read from file
	fileData, err := os.ReadFile(filename)
	if err != nil {
		// Log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// convert []byte to string
	bs := string(fileData)
	// convert string to []string
	s := strings.Split(bs, ",")
	// convert []string to deck
	d := deck(s)

	// d := deck(strings.Split(string(fileData), ","))
	return d
}
