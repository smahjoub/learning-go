package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length of 52 but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected 'Ace of Spades' but got '%v'", d[0])
	}

	if d[len(d)-1] != "Ten of Clubs" {
		t.Errorf("Expected 'Ten of Clubs' but got '%v'", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected length of 52 but got %d", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
