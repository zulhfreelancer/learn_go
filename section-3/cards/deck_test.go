package main

import "testing"

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
