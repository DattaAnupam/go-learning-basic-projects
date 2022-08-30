package main

import "testing"

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
