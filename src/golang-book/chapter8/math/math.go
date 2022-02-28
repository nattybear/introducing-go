package math

func Average(xs []float64) float64 {
  if len(xs) == 0 {
    return 0
  }
  total := float64(0)
  for _, x := range xs {
    total += x
  }
  return total / float64(len(xs))
}

func Min(xs []float64) float64 {
  min := xs[0]
  for _, x := range xs {
    if x < min {
      min = x
    }
  }
  return min
}

func Max(xs []float64) float64 {
  max := xs[0]
  for _, x := range xs {
    if x > max {
      max = x
    }
  }
  return max
}
