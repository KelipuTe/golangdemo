package main

import (
  "fmt"
  "runtime"
)

func main() {
  var ms1 runtime.MemStats
  runtime.ReadMemStats(&ms1) //读取内存使用情况
  fmt.Printf("%+v\n", ms1)

  var num [10]int = [10]int{}
  fmt.Println(num)

  var ms2 runtime.MemStats
  runtime.ReadMemStats(&ms2)
  fmt.Printf("%+v\n", ms2)

  fmt.Printf("%d\n", ms2.TotalAlloc-ms1.TotalAlloc)
}
