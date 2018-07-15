package main

import "fmt"

func main() {
  colors := map[string]string{
    "white": "#ffffff",
    "black": "#000000",
    "yellow": "#ffff00",
  }

  printMap(colors)
}

func printMap(c map[string]string)  {
  for key, value := range c {
    fmt.Println("Hex code for", key, "is", value)
  }
}
