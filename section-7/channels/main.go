package main

import (
  "fmt"
  "net/http"
)

func main() {
  links := []string {
    "http://google.com",
    "http://facebook.com",
    "http://amazon.com",
  }

  for _, link := range links {
    checkLink(link)
  }
}

func checkLink(link string) {
  // We don't care about response, we care about error here.
  _, err := http.Get(link)
  if err != nil {
    fmt.Println(link, "might be down")
    return // that's all, no more code to execute
  }
  fmt.Println(link, "is up!")
}
