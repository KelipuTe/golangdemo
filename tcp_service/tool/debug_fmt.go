package tool

import (
	"demo_golang/net_service/config"
	"fmt"
)

func DebugPrintln(a ...interface{}) {
  if 1 == config.APP_DEBUG {
    fmt.Printf("[debug]:")
    fmt.Println(a...)
  }
}
