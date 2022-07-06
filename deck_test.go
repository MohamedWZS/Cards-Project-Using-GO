package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := new_deck()

	if len(d) != 16 {
		t.Errorf("expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("expected first card to be Ace of Hearts, but got %v", d[0])
	}

	if d[len(d)-1] != "Three of Diamonds" {
		t.Errorf("expected last card to be Three of Diamonds, but got %v", d[len(d)-1])
	}
}
