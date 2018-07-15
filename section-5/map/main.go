package main

import "fmt"

func main() {
  // First way of writing map.
  // colors := map[string]string{
  //   "white": "#fff",
  //   "black": "#000",
  // }

  // Second way of writing map.
  // var colors map[string]string

  // Third way of writing map.
  colors := make(map[string]string)
  // Add key pair to the map.
  colors["white"] = "#fff"
  // Delete key pair from the map.
  delete(colors, "white")

  fmt.Println(colors)
}
