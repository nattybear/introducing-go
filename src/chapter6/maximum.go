package main

import "fmt"

func max(x, y int) int {
  if x > y {
    return x
  } else {
    return y
  }
}

func maximum(xs ...int) (int) {
  if len(xs) == 0 {
    panic("empty list")
  }
  if len(xs) == 1 {
    return xs[0]
  } else {
    return max(xs[0], maximum(xs[1:]...))
  }
}

func main() {
  xs := []int{1,2,3}
  fmt.Println(maximum(xs...))
}
