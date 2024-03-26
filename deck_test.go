package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but go %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first item to be `Ace of Spades` but got `%v`", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected the first item to be `Four of Clubs` but got `%v`", d[len(d)-1])
	}
}

func TestSaveNewDeckToFile(t *testing.T) {
	const deckSaveFilename = "__decktesting"
	os.Remove(deckSaveFilename)

	deck := newDeck()

	deck.saveToFile(deckSaveFilename)

	loadSavedDeck := newDeckFromFile(deckSaveFilename)

	if len(loadSavedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadSavedDeck))
	}

	os.Remove(deckSaveFilename)
}
