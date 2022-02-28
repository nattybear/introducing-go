package math

import "testing"

type testpair struct {
  values []float64
  min float64
}

var tests = []testpair{
    { []float64{1,2}, 1 },
    { []float64{1,1,1,1,1,1}, 1 },
    { []float64{-1,1}, -1 },
}

func TestAverage(t *testing.T) {
  for _, pair := range tests {
    v := Min(pair.values)
    if v != pair.min {
      t.Error(
        "For", pair.values,
        "expected", pair.min,
        "got", v,
      )
    }
  }
}
