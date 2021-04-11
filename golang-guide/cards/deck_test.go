package main

import (
	"os"
	"testing")


func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected 'Ace of Spades', nut got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected 'Four of Clubs', but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	file := "_decktesting"
	os.Remove(file)

	deck := newDeck()
	deck.save(file)

	loadedDeck := newDeckFromFile(file)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards, but got %v", len(loadedDeck))
	}

	os.Remove(file)
}