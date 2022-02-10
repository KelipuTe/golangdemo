package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println(time.Parse("2005-01-02", "2021-12-01"))
  fmt.Println(time.Parse("2006-01-02", "2016-07-08"))
}
