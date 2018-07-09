package main

import (
  "fmt"
  "strings"
  "io/ioutil"
  "os"
  "math/rand"
  "time"
)

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

// Helper 'method' to convert deck to string (comma separated).
func (d deck) toString() string {
  // convert deck to slice of strings
  sliceOfStrings := []string(d)
  // join all elements in that slice of strings
  return strings.Join(sliceOfStrings, ",")
}

// Save a file to disk.
// The 'error' is a type. The 'WriteFile' returns 'error',
// so we need to implement it here, just in case something is wrong.
// The 0666 means anyone can read and write this file.
func (d deck) saveToFile(filename string) error {
  deckInString := d.toString()
  sliceOfBytes := []byte(deckInString)
  // we 'return' here so that it will return 'error' (if we get one)
  return ioutil.WriteFile(filename, sliceOfBytes, 0666)
}

// Read from a file on the disk.
func newDeckFromFile(filename string) deck {
  sliceOfBytes, err := ioutil.ReadFile(filename)

  if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1) // non-zero means exit and the task was not successful
  }

  // convert slice of bytes to string
  byteSliceInString := string(sliceOfBytes)
  // remove the comma separators in the string
  cleanedString := strings.Split(byteSliceInString, ",")
  // convert string to deck (slice of strings)
  return deck(cleanedString)
}

// Randomize cards.
// Go 'rand#Intn' always starts with same source/seed.
// That means, the randomizer function always returns same number.
// If we change the source/seed, we will get different number.
// So here, we are relying on program's 'start at' time as the source/seed.
func (d deck) shuffle() {
  // everytime we start our Go program, we will different time
  unixNano := time.Now().UnixNano()
  // create a Rand source with 'unixNano' as the seed
  source := rand.NewSource(unixNano)
  // create a Rand object with that source
  r := rand.New(source)

  for i := range d {
    // generate a random number from 0 to end of slice (length - 1)
    newPosition := r.Intn(len(d) - 1)
    // for EACH iteration, swap the position of 'i' and 'newPosition'
    d[i], d[newPosition] = d[newPosition], d[i]
  }
}
