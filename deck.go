package main

import "fmt"

// create a new type 'deck'
// which is a slice of strings
type deck []string

// creating a receiver
func (d deck) print() {
	for i, cards := range d {
		fmt.Println(i, cards)
	}
}
