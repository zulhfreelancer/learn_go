package main

import "fmt"

type person struct {
  firstName string
  lastName string
}

func main() {
  var alex person
  alex.firstName = "Alex"
  alex.lastName = "Anderson"

  // will show {} if we didn't set the firstName and lastName
  fmt.Println(alex)

  // will show {firstName: lastName:} if we didn't set them
  fmt.Printf("%+v", alex)
}
