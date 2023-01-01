package main

import "fmt"

func main() {
	cards := newDeck()
	drawCards, cards := draw(5, cards)

	drawCards.saveToFile("testFile.txt")

	handCards := newDeckFromFile("testFile.txt")
	fmt.Println(handCards)
}
