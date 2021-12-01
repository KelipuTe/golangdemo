package main

import "fmt"

type handle func(x, y int)

func (h handle) test(x, y int) {
  h(x, y)
}

func main() {
  var a handle = func(x, y int) {
    fmt.Println(x, y)
  }
  a.test(1, 2)
}
