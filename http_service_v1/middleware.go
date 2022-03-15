package http_service_v1

import (
  "fmt"
  "time"
)

// MiddlewareFunc 中间件处理函数。
// 这里和 HTTPHandlerFunc 保持一致。要不然最后一环，调用没法串起来。
type MiddlewareFunc func(c *Context)

// MiddlewareBuilder 中间件建造器。
// 实现思路就是链式套娃。返回一个 MiddlewareFunc 函数。
// 在返回的函数内部会调用传入的 next MiddlewareFunc 函数。
type MiddlewareBuilder func(next MiddlewareFunc) MiddlewareFunc

func TestMiddlewareBuilder(next MiddlewareFunc) MiddlewareFunc {
  return func(c *Context) {
    fmt.Printf("request before test middleware.\n")
    next(c)
    fmt.Printf("request after test middleware.\n")
  }
}

func TimeCostMiddlewareBuilder(next MiddlewareFunc) MiddlewareFunc {
  return func(c *Context) {
    // 执行前的时间
    startTime := time.Now().UnixNano()
    next(c)
    // 执行后的时间
    endTime := time.Now().UnixNano()
    fmt.Printf("request time cost: %d unix nano.\n", endTime-startTime)
  }
}
