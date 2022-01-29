## Exercises

1. What does the following program print?

```go
i := 10
if i > 10 {
  fmt.Println("Big")
} else {
  fmt.Println("Small")
}
```

```
Small
```

2. Write a program that prints out all the numbers between 1 and 100 that are evenly divisible by 3 (i.e., 3,6,9,etc).

```go
package main

import "fmt"

func main() {
  for i := 1; i <= 100; i++ {
    if i % 3 == 0 {
      fmt.Println(i)
    }
  }
}
```

3. Write a program that prints the numbers from 1 to 100, but for multiples of three, print "Fizz" instead of the number, and for the multiples of five, print "Buzz". For numbers that are multiples of both three and five, print "FizzBuzz".

```go
package main

import "fmt"

func main() {
  for i := 1; i <= 100; i++ {
    if (i % 3 == 0 && i % 5 == 0) {
      fmt.Println("FizzBuzz")
    } else if i % 3 == 0 {
      fmt.Println("Fizz")
    } else if i % 5 == 0 {
      fmt.Println("Buzz")
    } else {
      fmt.Println(i)
    }
  }
}
```
