package main

func main() {
	// cards := newDeck()
	// cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// println()
	// remainingCards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("fisier")
	cards := newDeckFromFile("fisier")
	cards.print()
}
