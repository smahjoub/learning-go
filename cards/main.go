package main

func main() {
	cards := newDeck()

	hand, remainingHand := deal(cards, 30)
	hand.print()
	remainingHand.print()

}

func newCard() string {
	return "Five of Diamonds"
}
