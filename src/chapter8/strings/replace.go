package main

import (
  "fmt"
  "strings"
)

func main() {
  // func Replace(s, old, new string, n int) string
  fmt.Println(strings.Replace("aaaa", "a", "b", 2))
  // => "bbaa"
}
