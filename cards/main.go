package main

func main() {
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Six of Clubs")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
