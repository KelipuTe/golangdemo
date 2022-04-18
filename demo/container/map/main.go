package main

import "fmt"

func main() {
  mapInit()
  mapKV()
  mapFor()
  mapExist()
  mapDelete()
}

func mapInit() {
  fmt.Printf("mapInit\r\n")

  // map 的底层实现是个结构体，/src/runtime/map.go
  // 只声明没有初始化的 map 是 nil，不能直接用
  var t1mapa map[string]string
  fmt.Printf("t1mapa: %+v,  len: %d\r\n", t1mapa, len(t1mapa))
}

func mapKV() {
  fmt.Printf("mapKV\r\n")

  // 初始化空的 map，然后添加数据
  t1mapa := make(map[string]string)
  t1mapa["a"] = "a"
  t1mapa["b"] = "b"
  fmt.Printf("t1arr1a: %+v,  len: %d\r\n", t1mapa, len(t1mapa))

  // 初始化的时候，可以赋值，然后也可以添加数据
  t1mapb := map[string]string{"a": "a"}
  t1mapb["b"] = "b"
  fmt.Printf("t1mapb: %+v,  len: %d\r\n", t1mapb, len(t1mapb))
}

func mapFor() {
  fmt.Printf("mapFor\r\n")

  t1mapa := map[string]string{"a": "a", "b": "b", "c": "c"}

  // 只有一个接收参数时，取出的是键名
  for k := range t1mapa {
    fmt.Printf("t1mapa: k: %s; ", k)
  }
  fmt.Printf("\r\n")

  // 只有两个接收参数时，取出的是键和值
  for k, v := range t1mapa {
    fmt.Printf("t1mapa: k: %s, v: %s; ", k, v)
  }
  fmt.Printf("\r\n")
}

func mapExist() {
  fmt.Printf("mapExist\r\n")

  t1mapa := map[string]string{"a": "a", "b": "b", "c": "c"}

  // 可以通过这种写法，判断 map 中某个键是否存在。
  // 如果存在 ok 为 true，v 就是值。
  if v1, ok1 := t1mapa["a"]; ok1 {
    fmt.Printf("v1: %s\r\n", v1)
  } else {
    fmt.Printf("not exist\r\n")
  }

  // 如果不存在 ok 为 false，v 没有被赋值，理论上应该是 0 值
  if v2, ok2 := t1mapa["d"]; ok2 {
    fmt.Printf("v2: %s\r\n", v2)
  } else {
    fmt.Printf("not exist\r\n")
  }
}

func mapDelete() {
  fmt.Printf("mapDelete\r\n")

  t1mapa := map[string]string{"a": "a", "b": "b", "c": "c"}
  fmt.Printf("t1arr1a: %+v\r\n", t1mapa)

  // 通过 delete() 方法可以删除 map 中的某个键值对
  delete(t1mapa, "b")
  fmt.Printf("t1arr1a: %+v\r\n", t1mapa)

  // 如果要清空 map，放着不管等 GC 比手动删除效率更高。
}
