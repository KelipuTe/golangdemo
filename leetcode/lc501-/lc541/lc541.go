package main

import (
  "fmt"
)

func main() {
  // fmt.Println(reverseStr("abcd", 2))
  // fmt.Println(reverseStr("abcdefg", 2))
  fmt.Println(reverseStr("1234567812345678abcdefg", 8))
}

//给定一个字符串s和一个整数k，从字符串开头算起，每2k个字符反转前k个字符。
//如果剩余字符少于k个，则将剩余字符全部反转。
//如果剩余字符小于2k但大于或等于k个，则反转前k个字符，其余字符保持原样。
//1<=s.length<=10^4;s仅由小写英文组成;1<=k<=10^4

//双指针

//541-反转字符串II
func reverseStr(s string, k int) string {
  var sLen int = len(s)
  var sli1S []byte = []byte(s)
  var indexZuo3, indexYou4 int = 0, k - 1 //每次需要反转的范围

  for indexYou4 < sLen {
    var tIndexZuo3, tIndexYou4 int = indexZuo3, indexYou4
    for tIndexZuo3 < tIndexYou4 {
      sli1S[tIndexZuo3], sli1S[tIndexYou4] = sli1S[tIndexYou4], sli1S[tIndexZuo3]
      tIndexZuo3++
      tIndexYou4--
    }
    indexZuo3 += k << 1
    indexYou4 += k << 1
  }

  //处理剩余字符少于k个的情况
  if indexYou4 >= sLen && indexZuo3 < sLen {
    var tIndexZuo3, tIndexYou4 int = indexZuo3, sLen - 1
    for tIndexZuo3 < tIndexYou4 {
      sli1S[tIndexZuo3], sli1S[tIndexYou4] = sli1S[tIndexYou4], sli1S[tIndexZuo3]
      tIndexZuo3++
      tIndexYou4--
    }
  }

  return string(sli1S)
}
