package helper

import "fmt"

func PrintlnInDebug(AppDebug int, a ...interface{}) {
  if 1 == AppDebug {
    fmt.Printf("[debug]:")
    fmt.Println(a...)
  }
}
