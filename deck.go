package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// newDeck(), returns a deck of cards
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

// printCard(), prints the card
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

// deckToString(), returns string representation of the deck
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

// saveToFile(), Saves deck of card into the hard drive, with 'filename' and default permissions
func (d deck) saveToFile(filename string) error {
	// convert the deck into string
	s := d.deckToString()
	// convert the string to []byte
	sb := []byte(s)
	// bs := []byte(d.deckToString())
	return os.WriteFile(filename, sb, 0666) // 0666: anyone can read or write to the file
}

// getFromFile(), returns deck, for the given filename from hard drive
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

// shuffle(), takes the deck of card and randomly shuffles it
func (d deck) shuffle() {
	// random seed value for rand.Intn()
	// every time run the program, get an unique int64 value, to get unique seed value
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	for i := range d {
		// get an random integer from 0 to (length of deck -1)
		newPosition := r.Intn(len(d) - 1)

		// swap the two values
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
