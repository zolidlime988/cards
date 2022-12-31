package main

import (
	"fmt"
)

func main() {
	cards := deck{"Five Of Diamonds", newCard()}
	fmt.Println(cards)
	cards = append(cards, "Six Of Spades")
	cards.print()
}

func newCard() string {
	return "Ace Of Spades"
}
