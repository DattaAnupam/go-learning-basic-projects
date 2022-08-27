package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.printCard()
	remainingCards.printCard()
}
