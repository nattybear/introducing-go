package main

import (
  "fmt"
  "strings"
)

func main() {
  // func Split(s, sep string) []string
  fmt.Println(strings.Split("a-b-c-d-e", "-"))
  // => []string{"a","b","c","d","e"}
}
