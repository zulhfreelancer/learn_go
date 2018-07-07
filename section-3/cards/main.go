package main

func main() {
  cards := deck{"Ace of Diamonds", newCard()}

  // pushing a new card into the slice
  cards = append(cards, "Six of Spades")

  // call receiver in deck.go
  cards.print()
}

func newCard() string {
  return "Five of Diamonds"
}
