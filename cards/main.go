package main

func main() {
	cards := newDeck()
	cards.saveToFile("mycards.txt")

	newDeck := newDeckFromFile("mycards.txt")
	newDeck.print()
}
