package main

import (
  "fmt"
  "net/http"
  "time"
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
    // Create an anonymous function and execute it straight away.
    // To reduce confusion, the `_link` is the `link` passed at the end of this function.
    go func(_link string) {
      // Pause it for 5 seconds.
      time.Sleep(5 * time.Second)

      // Execute the next child routine.
      checkLink(_link, c)

      // Execute this anonymous function straight away in form of Go Routine,
      // passing `link` (from the `range` line) as argument.
      // That is same as calling a `nonAnonymousFunction(link)`.
    }(link)
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
