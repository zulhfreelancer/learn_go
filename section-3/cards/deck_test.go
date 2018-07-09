package main

import (
  "testing"
  "os"
)

// The 't' is the test handler
func TestNewDeck(t *testing.T) {
  d := newDeck()

  if len(d) != 16 {
    // The '%v' will take and render the value of next argument
    t.Errorf("Expected deck length of 16, but got %v", len(d))
  }

  if d[0] != "Ace of Spades" {
    t.Errorf("Expected first card is Ace of Spades, but got %v", d[0])
  }

  if d[len(d) - 1] != "Four of Clubs" {
    t.Errorf("Expected first card is Four of Clubs, but got %v", d[len(d) - 1])
  }
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
  filename := "_deck_testing.txt"

  // delete the old file
  os.Remove(filename)

  deck := newDeck()
  deck.saveToFile(filename)

  loadedDeck := newDeckFromFile(filename)

  if len(loadedDeck) != 16 {
    t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
  }

  os.Remove(filename)
}
