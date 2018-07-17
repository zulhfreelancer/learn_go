package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
)

type logWriter struct{}

func main() {
  resp, err := http.Get("https://www.google.com/")

  if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
  }

  lw := logWriter{}

  // Copy source (src) to destination (dst).
  // In other words, copy response body to stdout (using fmt.Println below).
  // Function signature:
  //   'func Copy(dst Writer, src Reader) (written int64, err error)'
  // Where...
  //   'lw' (destination) implements Writer, and
  //   'resp.Body' (source) implements Reader
  io.Copy(lw, resp.Body)
}

func (lw logWriter) Write(bs []byte) (int, error) {
  // print the HTTP response to stdout
  fmt.Println(string(bs))

  // print additional info to stdout
  fmt.Println("\nJust wrote this many bytes:", len(bs))

  // to satisfy the function signature of Writer
  return len(bs), nil
}
