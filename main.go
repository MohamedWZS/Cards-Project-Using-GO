package main

func main() {
	cards := deck{"ten of spades", "six of hearts"}
	cards = append(cards, "five of hearts")

	cards.print()
}
