package main

import "fmt"

func main() {
  // {} means it's a slice (an array that can shrink/grow)
  cards := []string{"Ace of Diamonds", newCard()}

  // pushing a new card into the slice
  cards = append(cards, "Six of Spades")

  for i, card := range cards {
    fmt.Println(i, card)
  }
}

func newCard() string {
  return "Five of Diamonds"
}
