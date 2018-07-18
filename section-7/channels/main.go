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

  // Create a channel.
  // We gonna share string in the channel.
  c := make(chan string)

  for _, link := range links {
    go checkLink(link, c)
  }

  // Infinite loop.
  // range line: Wait for channel (`c`) to return something and assign it to `link`.
  // go checkLink line: then run child routine, passing the link that we've just received
  // This loop syntax is equivalent with previous loop syntax.
  for link := range c {
    go checkLink(link, c)
  }
}

func checkLink(link string, c chan string) {
  // We don't care about response (first argument), we care about error here.
  _, err := http.Get(link)

  if err != nil {
    fmt.Println(link, "might be down")
    c <- link
    return // that's all, no more code to execute
  }

  fmt.Println(link, "is up!")
  c <- link
}
