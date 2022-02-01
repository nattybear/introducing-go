package main

import "fmt"

func f2() (r int) {
  r = 1
  return
}

func main() {
  fmt.Println(f2())
}
