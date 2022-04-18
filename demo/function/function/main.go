package main

import (
  "fmt"
)

func main() {
  fmt.Println(funcReturn1("funcReturn1"))
  fmt.Println(funcReturn2(1, "funcReturn2"))
  fmt.Println(funcReturn1WithName("funcReturn1WithName"))
  fmt.Println(funcReturn2WithName(1, "funcReturn2WithName"))
  fmt.Println(funcSameType(1, 2, "funcSameType"))
  funcUnknowLen(1, "a", "b", "c")
}

// funcReturn1 一个返回值，不需要括号括起来
func funcReturn1(t1string string) string {
  return "hello, " + t1string
}

// funcReturn2 多个返回值，需要括号括起来
func funcReturn2(t1int int, t1string string) (int, string) {
  return t1int + 1, "hello, " + t1string
}

// funcReturn1WithName 一个有名字的返回值
// 直接 return 返回的就是有名字的这个变量，也可以 return 其他的变量
func funcReturn1WithName(t1string string) (t1return string) {
  t1return = "hello, " + t1string
  return
}

// funcReturn2WithName 多个有名字的返回值
func funcReturn2WithName(t1int int, t1string string) (ret1int int, ret1string string) {
  return t1int + 1, "hello, " + t1string
}

// funcSameType 多个相同类型的参数，放在一起可以只写一次类型
func funcSameType(t1inta, t1intb int, t1string string) (ret1inta, ret1intb int, ret1string string) {
  ret1inta = t1inta + 1
  ret1intb = t1intb + 1
  ret1string = "hello, " + t1string
  return
}

// funcUnknowLen 有不定参数，不定参数要放在最后面
func funcUnknowLen(t1int int, t1arr1string ...string) {
  fmt.Println("funcClosure")

  // 使用的时候，可以直接把 t1arr1string 看做切片
  fmt.Printf("t1int: %d, t1arr1string: ", t1int)
  for _, k := range t1arr1string {
    fmt.Printf("%s, ", k)
  }
  fmt.Printf("\r\n")
}
