## Exercises

1. `sum` is a function that takes a slice of numbers and adds them together. What would its function signature look like in Go?

```go
func sum(args ...int) int
```

2. Write a function that takes an integer and halves it and returns true if it was even of false if it was odd. For example, `half(1)` should return `(0, false)` and `half(2)` should return `(1, true)`.

```go
func half(n int) (float64, bool) {
  m := float64(n / 2)
  return m, n % 2 == 0
}
```

3. Write a function with one variadic parameter that finds the greatest number in a list of numbers.

```go
func maximum(xs ...int) (int) {
  var max int
  for i, v := range xs {
    if i == 0 || v > max {
      max = v
    }
  }
  return max
}
```

4. Using `makeEvenGenerator` as an example, write a `makeOddGenerator` function that generates odd numbers.

```go
func makeOddGenerator() func() uint {
  i := uint(1)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}
```

5. The Fibonacci sequence is defined as: `fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2)`. Write a recursive function that can find `fib(n)`.

```go
func fib(n int) int {
  if n == 0 {
    return 0
  }
  if n == 1 {
    return 1
  }
  return fib(n-1) + fib(n-2)
}
```

6. What are `defer`, `panic`, and `recover`? How do you recover from a runtime panic?

* `defer` : special statement that schedules a function call to be run after the function completes
* `panic` : function to cause a runtime error
* `recover` : stops the `panic` and returns the value that was passwd to the call to `panic`

```go
func main() {
  defer func() {
    str := recover()
    fmt.Println(str)
  }()
  panic("PANIC")
}
```

7. How do you get the memory address of a variable?
```go
&xPtr
```

8. How do you assign a value to a pointer?

```go
*xPtr = 1
```

9. How do you create a new pointer?

```go
xPtr := new(int)
```

10. What is the value of `x` after running this program:

```go
func square(x *float64) {
  *x = *x * *x
}

func main() {
  x := 1.5
  square(&x)
}
```

```
2.25
```

11. Write a program that can swap two integers `(x := 1; y := 2; swap(&x, &y)` should give you `x=2` and `y=1`).

```go
func swap(x, y *int) {
  *x, *y = *y, *x
}
```

