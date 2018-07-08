package main

// import "fmt"

func main() {
  /* -----------------
  * Save to disk
  ----------------- */
  // cards := newDeck()
  // cards.saveToFile("myCards.txt")


  /* -----------------
  * Read from disk
  ----------------- */
  cards := newDeckFromFile("myCards.txt")
  cards.print()
}
