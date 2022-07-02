package main

func main() {
	filename := "cards.save"

	newDeck().shuffle().saveToFile(filename)

	deck := newDeckFromFile(filename)

	hand, _ := deal(deck, 5)

	hand.print()
}
