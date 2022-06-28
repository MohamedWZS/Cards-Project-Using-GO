package main

func main() {
	cards := new_deck()

	hand, reamainingCards := deal(cards, 5)

	hand.print()
	reamainingCards.print()
	cards.print()
}
