package main

import (
  "fmt"
  "strings"
  "unicode/utf8"
)

func main() {
  stringSingle()
  stringMultiple()
  stringUnicode()
  stringJoin()
  stringSplit()
}

// stringSingle 字符串
func stringSingle() {
  t1string := "hello, world"

  fmt.Printf("t1string: %s\r\n", t1string)

  // 计算字符串长度
  fmt.Printf("len(t1string): %d\r\n", len(t1string))

  // 字符串可以通过下标读取，但是不能通过下标修改
  fmt.Printf("t1string for: ")
  for i := 0; i < len(t1string); i++ {
    fmt.Printf("%c-", t1string[i])
  }
  fmt.Printf("\r\n")

  // 字符串不能直接修改，需要先转换成字符切片，改完之后再转换回来
  t1arr1t1string := []byte(t1string)
  t1arr1t1string[0] = 'i'
  t1string = string(t1arr1t1string)
  fmt.Printf("t1string: %s\r\n", t1string)
}

// stringMultiple 多行字符串
func stringMultiple() {
  // 通过反引号（`）声明，而且在多行字符串中转义字符无效。
  t1string := `hello, world\r\n
hello, world\r\n`

  fmt.Printf("t1string: \r\n%s\r\n", t1string)
}

// stringUnicode Unicode 字符串
func stringUnicode() {
  t1string := "你好，世界"

  fmt.Printf("t1string: %s\r\n", t1string)

  // Unicode 字符串不能用 len() 计算字符串长度，要用 utf8.RuneCountInString()
  fmt.Printf("len(t1string): %d\r\n", len(t1string))
  fmt.Printf("utf8.RuneCountInString(t1string): %d\r\n", utf8.RuneCountInString(t1string))

  // Unicode 字符串不能通过下标读取
  fmt.Printf("t1string for: ")
  for _, s := range t1string {
    fmt.Printf("%c-", s)
  }
  fmt.Printf("\r\n")
}

// stringJoin 字符串拼接
func stringJoin() {
  t1string, t2string := "hello", "world"

  // 通过加号(+)拼接
  t3string := t1string + ", " + t2string
  fmt.Printf("t3string: %s\r\n", t3string)

  // 通过 fmt.Sprintf() 格式化输出，变相达到拼接的目的
  t4string := fmt.Sprintf("%s, %s", t1string, t2string)
  fmt.Printf("t4string: %s\r\n", t4string)
}

// stringSplit 字符串截取
func stringSplit() {
  t1string := "hello, world"
  // 先找到截断的位置
  index := strings.Index(t1string, ",")
  // 然后用切片操作。注意，切片操作左闭右开
  t2string := t1string[:index]
  fmt.Printf("t2string: %s\r\n", t2string)
  t3string := t1string[index:]
  fmt.Printf("t3string: %s\r\n", t3string)
}
