package main

import (
  "fmt"
  "math"
  "unsafe"
)

type StructType struct {
  t1int int
  p1int *int
}

type FuncType func(t1int int)

func main() {
  varInit()
  varExchange()
  varExchangeInt()
}

// varInit 变量初始化
// 声明变量时，会自动对变量对应的内存区域进行初始化操作。
// 初始化也就是赋 0 值。所以需要额外注意 0 值有效的情况。
func varInit() {
  var (
    t1bool bool

    t1string string

    t1int     int
    t1int32   int32
    t1int64   int64
    t1float32 float32
    t1float64 float64

    t1p1int *int
    t1func  FuncType

    t1arr1array [2]int
    t1arr1slice []int
    t1map       map[string]string

    t1struct StructType
  )

  fmt.Printf("t1bool, type: %T, value: %#v, sizeof: %d\r\n", t1bool, t1bool, unsafe.Sizeof(t1bool))

  fmt.Printf("t1string, type: %T, value: %#v, sizeof: %d\r\n", t1string, t1string, unsafe.Sizeof(t1string))

  fmt.Printf("t1int, type: %T, value: %#v, sizeof: %d\r\n", t1int, t1int, unsafe.Sizeof(t1int))
  fmt.Printf("t1int32, type: %T, value: %#v, sizeof: %d\r\n", t1int32, t1int32, unsafe.Sizeof(t1int32))
  fmt.Printf("t1int64, type: %T, value: %#v, sizeof: %d\r\n", t1int64, t1int64, unsafe.Sizeof(t1int64))
  fmt.Printf("t1float32, type: %T, value: %#v, sizeof: %d\r\n", t1float32, t1float32, unsafe.Sizeof(t1float32))
  fmt.Printf("t1float64, type: %T, value: %#v, sizeof: %d\r\n", t1float64, t1float64, unsafe.Sizeof(t1float64))

  fmt.Printf("math.MaxInt: %d\r\n", math.MaxInt)
  fmt.Printf("math.MaxInt32: %d\r\n", math.MaxInt32)
  fmt.Printf("math.MaxInt64: %d\r\n", math.MaxInt64)
  fmt.Printf("math.MaxFloat32: %f\r\n", math.MaxFloat32)
  fmt.Printf("math.MaxFloat64: %f\r\n", math.MaxFloat64)

  fmt.Printf("t1p1int, type: %T, value: %#v, sizeof: %d\r\n", t1p1int, t1p1int, unsafe.Sizeof(t1p1int))
  fmt.Printf("t1func, type: %T, value: %#v, sizeof: %d\r\n", t1func, t1func, unsafe.Sizeof(t1func))

  fmt.Printf("t1arr1array, type: %T, value: %#v, sizeof: %d\r\n", t1arr1array, t1arr1array, unsafe.Sizeof(t1arr1array))
  fmt.Printf("t1arr1slice, type: %T, value: %#v, sizeof: %d\r\n", t1arr1slice, t1arr1slice, unsafe.Sizeof(t1arr1slice))
  fmt.Printf("t1map, type: %T, value: %#v, sizeof: %d\r\n", t1map, t1map, unsafe.Sizeof(t1map))

  fmt.Printf("t1struct, type: %T, value: %#v, sizeof: %d\r\n", t1struct, t1struct, unsafe.Sizeof(t1struct))
}

// varExchange 交换两个变量
func varExchange() {
  // 在 go 里可以这么写
  t1c, t1d := 100, 200
  fmt.Printf("t1c: %d, t1d: %d\r\n", t1c, t1d)
  t1c, t1d = t1d, t1c
  fmt.Printf("t1c: %d, t1d: %d\r\n", t1c, t1d)
}

// varExchangeInt 交换两个 int 型变量
func varExchangeInt() {
  // 不用中间变量的方法
  t1a, t1b := 100, 200
  fmt.Printf("t1a: %d, t1b: %d\r\n", t1a, t1b)
  fmt.Printf("t1a: %b, t1b: %b\r\n", t1a, t1b)
  t1a = t1a ^ t1b
  fmt.Printf("t1a: %b, t1b: %b\r\n", t1a, t1b)
  t1b = t1b ^ t1a
  fmt.Printf("t1a: %b, t1b: %b\r\n", t1a, t1b)
  t1a = t1a ^ t1b
  fmt.Printf("t1a: %b, t1b: %b\r\n", t1a, t1b)
  fmt.Printf("t1a: %d, t1b: %d\r\n", t1a, t1b)
}
