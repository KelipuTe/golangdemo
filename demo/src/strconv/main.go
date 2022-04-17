package main

import (
  "fmt"
  "strconv"
)

func main() {
  intToString()
  stringToInt()
  stringParseInt()
  stringParseFloat()
}

// intToString int 转 string
func intToString() {
  t1int := 100
  t1string := strconv.Itoa(t1int)
  fmt.Printf("t1string:%T t1string:%s\n", t1string, t1string)
}

// stringToInt string 转 int
func stringToInt() {
  t1string := "200"
  t1int, err := strconv.Atoi(t1string)
  fmt.Printf("t1int:%T, t1int:%d, err: %v\n", t1int, t1int, err)
}

// stringParseInt string 转 int
func stringParseInt() {
  t1string := "300"
  // strconv.ParseInt(s string, base int, bitSize int)
  // base，10 表示转换成 10 进制；bitSize，64 表示转换成 int64 类型；
  t1int, err := strconv.ParseInt(t1string, 10, 64)
  fmt.Printf("t1int:%T, t1int:%d, err: %v\n", t1int, t1int, err)
}

// stringParseFloat string 转 float
func stringParseFloat() {
  t1string := "3.14"
  // strconv.ParseFloat(s string, bitSize int)
  // bitSize，64 表示转换成 int64 类型；
  t1float, err := strconv.ParseFloat(t1string, 64)
  fmt.Printf("t1float:%T, t1float:%f, err: %v\n", t1float, t1float, err)
}
