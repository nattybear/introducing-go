package main

import "fmt"

func half(n int) (float64, bool) {
  m := float64(n / 2)
  return m, n % 2 == 0
}

func main() {
  var n int
  fmt.Scan(&n)
  fmt.Println(half(n))
}
