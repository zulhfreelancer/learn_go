package main

import (
  "fmt"
  "net/http"
  "os"
)

func main() {
  resp, err := http.Get("https://www.google.com/")

  if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
  }

  // make a byte slice with 99,999 seats inside it
  bs := make([]byte, 99999)
  resp.Body.Read(bs)
  // convert to string, otherwise we just gonna see numbers e.g. [60 33 100 ...]
  fmt.Println(string(bs))
}
