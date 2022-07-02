package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck legth of 16, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be an Ace of Spades, but got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card to be an Four of Clubs, but got %s", d[0])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)
	
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck legth of 16, but got %d", len(loadedDeck))
	}

	os.Remove(filename)
}
