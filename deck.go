package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a type deck which is slice of strings
type deck []string

// functions
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func (d deck) shuffleDeck(times int) deck {
	if times == 0 {
		fmt.Printf("Shuffle amount remain %d. Return Deck\n", times)
		return d
	}

	// seed random number
	rand.Seed(time.Now().UnixMicro())

	// shuffle by seed
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
	fmt.Printf("Shuffle amount remain %d\n", times)

	// recursive shuffle
	return d.shuffleDeck(times - 1)
}

func (d deck) toString() string {
	return strings.Join(d, "\n")
}

func (d deck) shuff() deck {
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)

	for i := range d {
		newInd := rand.Intn(len(d) - 1)
		d[i], d[newInd] = d[newInd], d[i]
	}
	return d
}

// recievers
func newDeck() deck {

	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cardType := []string{"Spades", "Clubs", "Hearts", "Diamonds"}

	newDeck := make([]string, 52)
	for i, v := range cardValues {
		for j, t := range cardType {
			newDeck[i+j*13] = v + " of " + t
		}
	}
	// // seed random number
	// rand.Seed(time.Now().UnixMicro())

	// // shuffle by seed
	// rand.Shuffle(len(newDeck), func(i, j int) {
	// 	newDeck[i], newDeck[j] = newDeck[j], newDeck[i]
	// })

	return newDeck
}
func draw(numberOfCards int, deckCards deck) (deck, deck) {
	return deckCards[0:numberOfCards], deckCards[numberOfCards:]
}
func newDeckFromFile(file string) deck {
	bs, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("error reading: %v\n", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), "\n"))
}

func (deck deck) saveToFile(file string) {
	err := os.WriteFile(file, byteString(strings.Join(deck, "\n")), 0666)
	if err != nil {
		fmt.Printf("error writing: %v\n", err)
		os.Exit(1)
	}
}
