package main

import "fmt"

func fib(n int) int {
  if n == 0 {
    return 0
  }
  if n == 1 {
    return 1
  }
  return fib(n-1) + fib(n-2)
}

func main() {
  var n int
  fmt.Scan(&n)
  fmt.Println(fib(n))
}
