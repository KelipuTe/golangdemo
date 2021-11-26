package main

import (
  "fmt"
)

//先defer的语句在后面执行，因为底层用的就是栈结构
func main() {
  defer fmt.Println("defer1")
  defer fmt.Println("defer2")
  defer fmt.Println("defer3")
}
