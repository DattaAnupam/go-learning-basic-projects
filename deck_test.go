package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// total no of cards 52, (13*4)
	if len(d) != 52 {
		t.Errorf("Expected deck of length 52, but got %d", len(d))
	}
}
