package main

import "fmt"

func main() {
  cards := deck{"Ace of Diamonds", newCard()}

  // pushing a new card into the slice
  cards = append(cards, "Six of Spades")

  for i, card := range cards {
    fmt.Println(i, card)
  }
}

func newCard() string {
  return "Five of Diamonds"
}
