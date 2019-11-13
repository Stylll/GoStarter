package main

func main() {
	var cards = newDeckFromFile("MyCards")

	cards.shuffle()
	cards.print()
	// hand, remainingCards := deal(cards, 6)

	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())

	// cards.saveToFile("MyCards")

}
