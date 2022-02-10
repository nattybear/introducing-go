package main

import "fmt"

type Person struct {
  Name string
}

type Android struct {
  Person
  Model string
}

func (p *Person) Talk() {
  fmt.Println("Hi, my name is", p.Name)
}

func main() {
  a := new(Android)
  a.Talk()
}
