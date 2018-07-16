package main

import "fmt"

type bot interface {
  getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
  eb := englishBot{}
  sb := spanishBot{}

  printGreeting(eb)
  printGreeting(sb)
}

func printGreeting(b bot) {
  fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
  // custom logic for englishBot goes here...
  return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
  // custom logic for spanishBot goes here...
  return "Hola!"
}
