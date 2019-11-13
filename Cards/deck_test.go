package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length to be 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFileName := "_decktesting"
	os.Remove(testFileName)

	d := newDeck()
	d.saveToFile(testFileName)

	loadedDeck := newDeckFromFile(testFileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length to be 16 but got %v", len(d))
	}

	os.Remove(testFileName)

}
