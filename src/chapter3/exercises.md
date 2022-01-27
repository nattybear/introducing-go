1. What are two ways to create a new variable?
```go
var
const
```

2. What is the value of `x` after running `x := 5; x += 1`?
```go
6
```

3. What is scope? How do you determine the scope of a variable in Go?
> The range of places where you are allowed to use `x` is called the scope of the variable. According to the language specification, "Go is lexically scoped using blocks."

4. What is the difference between `var` and `const`?
> Constants are essentially variables whose values cannot be changed later.

5. Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius `(C = (F - 32) * 5/9)`.
[87f5eff](https://github.com/nattybear/introducing-go/commit/87f5effaa6793cd284c66af8bf13623da519436a)

6. Write another program that converts from feet into meters (`1ft = 0.3048m`).
[9690dc9](https://github.com/nattybear/introducing-go/commit/9690dc98268e5446841fe3f535afafa4d3baf929)
