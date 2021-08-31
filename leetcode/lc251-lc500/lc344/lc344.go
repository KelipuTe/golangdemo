package main

import "fmt"

func main() {
  var sli1Chars []byte = []byte{'h', 'e', 'l', 'l', 'o'}
  fmt.Println(sli1Chars)
  reverseString(sli1Chars)
  fmt.Println(sli1Chars)
}

//编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组char[]的形式给出。
//不要给另外的数组分配额外的空间，必须原地修改输入数组、使用O(1)的额外空间解决这一问题。
//可以假设数组中的所有字符都是ASCII码表中的可打印字符。

//344-反转字符串
func reverseString(s []byte) {
  var sLen int = len(s)

  for index := 0; index < sLen>>1; index++ {
    s[index], s[sLen-1-index] = s[sLen-1-index], s[index]
  }
}
