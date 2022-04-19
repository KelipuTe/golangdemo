package main

import (
  "fmt"
  "unicode"
)

func main() {
  // Go 支持 Unicode（UTF-8）。Unicode 与 ASCII 类似，都是一种字符集。
  // 字符集为每个字符分配一个唯一的 ID，所有字符在 Unicode 字符集中都有一个唯一的 ID。
  // UTF-8 是编码规则，将 Unicode 中字符的 ID 以某种方式进行编码

  var t1byte byte = 'h'
  fmt.Printf("t1byte: %c, asicc: %d, unicode: %d, utf-8: %U\r\n", t1byte, t1byte, t1byte, t1byte)

  var t1rune rune = '你'
  fmt.Printf("t1rune: %c, unicode: %d, utf-8: %U\r\n", t1rune, t1rune, t1rune)

  // unicode 包提供了一些测试字符类型的方法
  var t2rune rune = 'a'
  fmt.Printf("unicode.IsLetter(t2rune): %t,\r\n", unicode.IsLetter(t2rune))
  fmt.Printf("unicode.IsDigit(t2rune): %t,\r\n", unicode.IsDigit(t2rune))

  var t3rune rune = '1'
  fmt.Printf("unicode.IsLetter(t3rune): %t,\r\n", unicode.IsLetter(t3rune))
  fmt.Printf("unicode.IsDigit(t3rune): %t,\r\n", unicode.IsDigit(t3rune))
}
