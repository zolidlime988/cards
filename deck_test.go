package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected 52, got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace Of Spades, got %v", d[0])
	}
	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected King of Diamond, got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testFileName := "_deckTesting.txt"
	os.Remove(testFileName)
	deck := newDeck()
	deck.saveToFile(testFileName)

	d := newDeckFromFile(testFileName)
	if len(d) != 52 {
		t.Errorf("Expected 52, got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace Of Spades, got %v", d[0])
	}
	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected King of Diamonds, got %v", d[len(d)-1])
	}
}
