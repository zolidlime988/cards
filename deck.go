package main

import (
	"fmt"
	"math/rand"
	"time"
)

// create a type deck which is slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {

	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cardType := []string{"Spades", "Clubs", "Hearts", "Diamonds"}

	newDeck := make([]string, 52)
	for i, v := range cardValues {
		for j, t := range cardType {
			newDeck[i+j*13] = v + " of " + t
		}
	}
	// seed random number
	rand.Seed(time.Now().UnixMicro())

	// shuffle by seed
	rand.Shuffle(len(newDeck), func(i, j int) {
		newDeck[i], newDeck[j] = newDeck[j], newDeck[i]
	})

	return newDeck
}

func (d deck) shuffle(times int) deck {
	for i := 0; i < times; i++ {
		// seed random number
		rand.Seed(time.Now().UnixMicro())

		// shuffle by seed
		rand.Shuffle(len(d), func(i, j int) {
			d[i], d[j] = d[j], d[i]
		})
		fmt.Printf("Shuffle amount %d of %d\n", i+1, times)
	}
	return d
}
