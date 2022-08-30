package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// total no of cards 52, (13*4)
	if len(d) != 52 {
		t.Errorf("Expected deck of length 52, but got %d", len(d))
	}

	// Check first card of the is Spades of Ace
	if d[0] != "Spades of Ace" {
		t.Errorf("Expected first card as Spades of Ace, but got %s", d[0])
	}

	// Check last card of the is Clubs of Queen
	if d[len(d)-1] != "Clubs of Queen" {
		t.Errorf("Expected last card as Clubs of Queen, but got %s", d[len(d)-1])
	}
}

func TestSaveToFileAndGetFromFile(t *testing.T) {
	expCardLen := 52
	// clean any previous left over dummy file of testing
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := getFromFile("_decktesting")

	if len(loadedDeck) != expCardLen {
		t.Errorf("Expected deck of length %d, but got %d", expCardLen, len(loadedDeck))
	}

	// clean dummy file of testing
	os.Remove("_decktesting")
}

func TestDeal(t *testing.T) {
	deck := newDeck()
	d1, d2 := deal(deck, 3)

	if len(d1) == 0 {
		t.Errorf("Expected element in the first deck, but got nothing")
	}

	if len(d2) == len(deck) {
		t.Errorf("Second deck should have less amount of cards than the original one")
	}
}
