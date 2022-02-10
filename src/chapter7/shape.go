package main

import ("fmt"; "math")

type Rectangle struct {
  x1, y1, x2, y2 float64
}

type Circle struct {
  x, y, r float64
}

type Shape interface {
  arae() float64
}

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

func totalArea(circles ...Circle) float64 {
  var total float64
  for _, c := range circles {
    total += c.area()
  }
  return total
}

func main() {
  c1 := Circle{0, 0, 5}
  c2 := Circle{0, 0, 6}
  fmt.Println(totalArea(c1, c2))
}
