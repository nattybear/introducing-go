package main

import (
  "fmt"
  "strings"
)

func main() {
  // func Join(a []string, sep string) string
  fmt.Println(strings.Join([]string{"a","b"}, "-"))
  // => "a-b"
}
