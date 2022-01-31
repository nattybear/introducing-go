## Exercises

1. How do you access the fourth element of an array or slice?

```go
x[3]
```

2. What is the length of a slice created using `make([]int, 3, 9)`?

```
3
```

3. Given the following array, what would `x[2:5]` give you?

```go
x := [6]string{"a","b","c","d","e","f"}
```

```go
["c","d","e"]
```

4. Write a program that finds the smallest number in this list:

```go
x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
}
```

```go
package main

import "fmt"

func main() {
  var min int
  x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
  }
  for i, v := range x {
    if i == 0 || v < min {
      min = v
    }
  }
  fmt.Println(min)
}
```
