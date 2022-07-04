package main

import (
	"fmt"
	"strings"
)

// create a new type 'deck'
// which is a slice of strings
type deck []string

func new_deck() deck {
	myNewDeck := deck{}

	cardSuits := []string{"Hearts", "Spades", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "One", "Two", "Three"}

	for _, value := range cardValues {
		for _, suit := range cardSuits {
			myNewDeck = append(myNewDeck, value+" of "+suit)
		}
	}

	return myNewDeck
}

func (d deck) print() {
	for i, cards := range d {
		fmt.Println(i, cards)
	}
}

func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}

func (d deck) toString() string {
	fmt.Println("converting a deck to a unified string")
	return strings.Join([]string(d), ", ")
}
