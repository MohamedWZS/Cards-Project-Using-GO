package main

import (
	"fmt"
	"log"
)

func main() {
	cards := new_deck()

	fmt.Println(cards.toString())

	err := cards.saveToFile("my_cards")

	if err != nil {
		log.Fatal(err)
	}
}
