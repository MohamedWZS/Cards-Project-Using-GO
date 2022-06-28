package main

import "fmt"

// create a new type 'deck'
// which is a slice of strings
type deck []string

// create a new deck
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

// print receiver func
func (d deck) print() {
	for i, cards := range d {
		fmt.Println(i, cards)
	}
}

// deal func
func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}
