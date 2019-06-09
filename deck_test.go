package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v instead", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card to be Ace of Hearts, but go %v instead", d[0])
	}

	if d[len(d)-1] != "Four of Diamonds" {
		t.Errorf("Expected first card to be Four of Diamonds, but go %v instead", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v instead", len(loadedDeck))
	}
}
