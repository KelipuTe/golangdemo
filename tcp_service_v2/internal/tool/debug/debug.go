package debug

import (
  "fmt"
)

func Println(isDebug bool, sli1data ...interface{}) {
  if isDebug {
    fmt.Printf("[debug]: ")
    fmt.Println(sli1data...)
  }
}
