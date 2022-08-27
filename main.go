package main

func main() {
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Spades")

	cards.printCard()
}

func newCard() string {
	return "Five of Diamonds"

}
