package main

import "fmt"

type Data struct {
  arr1a []int      // 切片
  ida   InnerData  // InnerData
  p1id  *InnerData // 指针
}

type InnerData struct {
  t1int    int
  t1string string
}

func main() {
  slicePass()
  sliceMap()
  valuePass()
}

// slicePass 切片是引用传递
func slicePass() {
  fmt.Printf("slicePass\r\n")
  arr1a := []int{1, 2, 3}
  fmt.Printf("arr1a, p: %p, v: %+v\r\n", arr1a, arr1a)
  arr1b := slicePass2(arr1a)
  fmt.Printf("arr1a, p: %p, v: %+v\r\n", arr1a, arr1a)
  fmt.Printf("arr1b, p: %p, v: %+v\r\n", arr1b, arr1b)
}

func slicePass2(arr1c []int) []int {
  fmt.Printf("arr1c, p: %p, v: %+v\r\n", arr1c, arr1c)
  arr1d := arr1c
  fmt.Printf("arr1d, p: %p, v: %+v\r\n", arr1d, arr1d)
  arr1d[1] = 0
  return arr1d
}

// sliceMap map 是引用传递
func sliceMap() {
  fmt.Printf("sliceMap\r\n")
  mapa := map[string]string{"a": "a", "b": "b", "c": "c"}
  fmt.Printf("mapa, p: %p, v: %+v\r\n", mapa, mapa)
  mapb := sliceMap2(mapa)
  fmt.Printf("mapa, p: %p, v: %+v\r\n", mapa, mapa)
  fmt.Printf("mapb, p: %p, v: %+v\r\n", mapb, mapb)
}

func sliceMap2(mapc map[string]string) map[string]string {
  fmt.Printf("mapc, p: %p, v: %+v\r\n", mapc, mapc)
  mapd := mapc
  fmt.Printf("mapd, p: %p, v: %+v\r\n", mapd, mapd)
  mapd["b"] = "bb"
  return mapd
}

func valuePass() {
  fmt.Printf("valuePass\r\n")
  t1da := Data{
    arr1a: []int{1, 2, 3},
    ida:   InnerData{11, "aa"},
    p1id:  &InnerData{22, "bb"},
  }
  fmt.Printf("t1da, p: %p, v: %+v, arr1a p: %p\r\n", &t1da, t1da, &t1da.arr1a)
  t1db := valuePass2(t1da)
  fmt.Printf("t1db, p: %p, v: %+v, arr1a p: %p\r\n", &t1db, t1db, &t1db.arr1a)
}
func valuePass2(t1dc Data) Data {
  fmt.Printf("t1dc, p: %p, v: %+v, arr1a p: %p\r\n", &t1dc, t1dc, &t1dc.arr1a)
  t1dd := t1dc
  t1dd.arr1a[1] = 0
  t1dd.ida.t1int = 22
  fmt.Printf("t1dd, p: %p, v: %+v, arr1a p: %p\r\n", &t1dd, t1dd, &t1dd.arr1a)
  return t1dd
}
