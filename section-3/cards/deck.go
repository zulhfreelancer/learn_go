package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings.
// In object oriented, it's like we are extending String class
// here into 'deck' class. Note: Go doesn't has class.
type deck []string

// Create receiver with 'print' as the name
func (d deck) print() {
  for i, card := range d {
    fmt.Println(i, card)
  }
}
