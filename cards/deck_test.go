package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	newDeck := newDeck()

	if len(newDeck) != 16 {
		t.Errorf("the length of the deck should be 16 but got %v", len(newDeck))
	}

	if newDeck[0] != "Ace of Clubs" {
		t.Errorf("the first card of the deck should be Ace of Clubs but got %v", newDeck[0])
	}

	if newDeck[len(newDeck)-1] != "Four of Diamomds" {
		t.Errorf("the last card of the deck should be Four of Diamonds but got %v", newDeck[len(newDeck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	d := newDeck()
	d.saveToFile("_deckTesting")

	loadedDeck, _ := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {
		t.Errorf("the length of the loadedDeck should be 16 but got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
