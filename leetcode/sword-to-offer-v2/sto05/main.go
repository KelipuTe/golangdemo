package main

import "fmt"

// 剑指 Offer 05. 替换空格
// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

func main() {
  fmt.Println(replaceSpace("We are happy."))
}

func replaceSpace(s string) string {
  sLen := len(s)
  sli1s := make([]byte, sLen*3)
  j := 0
  for i := 0; i < sLen; i++ {
    if ' ' == s[i] {
      sli1s[j] = '%'
      sli1s[j+1] = '2'
      sli1s[j+2] = '0'
      j = j + 3
    } else {
      sli1s[j] = s[i]
      j = j + 1
    }
  }
  return string(sli1s[:j])
}
