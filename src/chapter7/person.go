package main

import "fmt"

type Person struct {
  Name string
}

func (p *Person) Talk() {
  fmt.Println("Hi, my name is", p.Name)
}

func main() {
  p := Person{"Vincent"}
  p.Talk()
}
