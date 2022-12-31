package main

import "fmt"

var card string

func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
