package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 , but got %v", len(d))
	}
	if d[0] != "Ace Of Spades" {
		t.Errorf("Expected first card to be Ace Of Spades , but found %v", d[0])
	}
}
