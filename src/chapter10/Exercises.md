## Exercises

### How do you specify the direction of a channel type?

```go
func pinger(c chan<- string)
```

> `pinger` is only allowed to send to `c`.

```go
func printer(c <- chan string)
```

### Write your own `Sleep` function using `time.After`.

### What is a buffered channel? How would you create one with a capacity of 20?

> A buffered channel is asynchronous; sending or receiving a message will not wait unless the channel is already full.

```go
c := make(chan int, 20)
```
