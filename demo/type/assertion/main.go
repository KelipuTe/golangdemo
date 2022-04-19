package main

import (
  "fmt"
)

func main() {
  getType(1)
  getType("a")
  getType([]int{1, 2})
}

func getType(a interface{}) {
  switch a.(type) {
  case int:
    fmt.Println("the type of a is int")
  case string:
    fmt.Println("the type of a is string")
  case float64:
    fmt.Println("the type of a is float")
  default:
    fmt.Println("unknown type")
  }
}
