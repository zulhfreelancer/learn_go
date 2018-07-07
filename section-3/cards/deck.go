package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings.
// In object oriented, it's like we are extending String class
// here into 'deck' class. Note: Go doesn't has class.
type deck []string

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
