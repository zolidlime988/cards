package main

func main() {
	cards := newDeck()
	drawCards, _ := draw(5, cards)

	drawCards.saveToFile("testFile.txt")

	handCards := newDeckFromFile("testFile.txt")
	handCards.shuff()
	handCards.print()
}
