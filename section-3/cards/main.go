package main

func main() {
  cards := newDeck()

  // call receiver in deck.go
  cards.print()
}
