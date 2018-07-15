package main

import "fmt"

type contactInfo struct {
  email string
  zipCode int
}

type person struct {
  firstName string
  lastName string
  contact contactInfo
}

func main() {
  jim := person{
    firstName: "Jim",
    lastName: "Party",
    contact: contactInfo{
      email: "jim@gmail.com",
      zipCode: 12345, // important to put comma here
    }, // important to put comma here
  }

  fmt.Printf("%+v", jim)
}
