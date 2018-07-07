package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings.
// In object oriented, it's like we are extending String class
// here into 'deck' class. Note: Go doesn't has class.
type deck []string

func newDeck() deck {
  cards := deck{}

  cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
  cardValues := []string{"Ace", "Two", "Three", "Four"}

  // The '_' tells Go that we know there is something at that place
  // but we don't care or don't want to use it. In this case, '_' is index.
  for _, suit := range cardSuits {
    for _, value := range cardValues {
      cards = append(cards, value + " of " + suit)
    }
  }

  return cards
}

// Create receiver with 'print' as the name.
// Any variables of type 'deck' in our project will get access
// to this 'print' method.
// The 'self' can be anything. By convention, it should be just 'd'.
// Note: there is no such thing called 'self' or 'this' in Go.
func (self deck) print() {
  for i, card := range self {
    fmt.Println(i, card)
  }
}

// This function returns 2 values and they are both 'deck' type.
func deal(d deck, handSize int) (deck, deck) {
  // First part takes from first element up to 'handSize' EXCLUDING 'handSize'.
  // Second part takes from 'handSize' to the last element.
  return d[:handSize], d[handSize:]
}
