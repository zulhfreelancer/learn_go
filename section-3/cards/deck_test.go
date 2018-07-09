package main

import "testing"

// The 't' is the test handler
func TestNewDeck(t *testing.T) {
  d := newDeck()

  if len(d) != 16 {
    // The '%v' will take and render the value of next argument
    t.Errorf("Expected deck length of 16, but got %v", len(d))
  }
}
