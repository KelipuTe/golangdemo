package main

import (
  "fmt"
  "math"
)

func main() {
  a := math.Floor(float64(11) / 3 * math.Pow10(2))
  a = a / math.Pow10(2)

  b := float64(11) - 3*a
  fmt.Println(b)
  // split(11, 3, 2)
}

func split(sum float64, len int, precision int) []float64 {
  sli1ret := make([]float64, 0, 2)

  return sli1ret
}
