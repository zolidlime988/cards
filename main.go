package main

import "fmt"

func main() {
	cards := newDeck()
	drawCards, cards := draw(5, cards)
	bString := byteString(drawCards.toString())

	err := writeToFile("testFile.txt", bString)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fileContext, err := getDataFromFile("testFile.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Println(string(fileContext))
}
