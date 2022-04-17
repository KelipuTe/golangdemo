package main

import (
  "errors"
  "fmt"

  "golang.org/x/sync/errgroup"
)

func main() {
  p1group := &errgroup.Group{}
  nums := []int{-1, 0, 1}
  for _, num := range nums {
    t1num := num
    p1group.Go(func() error {
      fmt.Println(t1num)
      if t1num < 0 {
        return errors.New("< 0")
      }
      return nil
    })
  }

  if err := p1group.Wait(); nil != err {
    fmt.Println("error=", err.Error())
  } else {
    fmt.Println("ok")
  }
}
