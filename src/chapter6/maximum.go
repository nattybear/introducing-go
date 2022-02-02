package main

import "fmt"

func maximum(xs ...int) (int) {
  var max int
  for i, v := range xs {
    if i == 0 || v > max {
      max = v
    }
  }
  return max
}

func main() {
  xs := []int{1,2,3}
  fmt.Println(maximum(xs...))
}
