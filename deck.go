package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	// read from file
	bs, err := ioutil.ReadFile(fileName)

	// error handling
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",")
	return deck(ss)
}
