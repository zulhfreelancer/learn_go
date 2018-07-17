package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
)

func main() {
  resp, err := http.Get("https://www.google.com/")

  if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
  }

  // func Copy(dst Writer, src Reader) (written int64, err error)
  // 'os.Stdout' implements Writer
  // while 'resp.Body' implements Reader
  io.Copy(os.Stdout, resp.Body)
}
