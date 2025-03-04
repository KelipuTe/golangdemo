package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(reverseVowels("hello"))
  fmt.Println(reverseVowels("leetcode"))
}

//双指针
//一个指针从前往后遍历，找到元音字母就停下，另一个指针从后往前遍历，找到元音字母就停下。
//交换两个指针指向的字母，然后两个指针各自向前进一步，继续遍历。

//345-反转字符串中的元音字母
func reverseVowels(s string) string {
  var sLen int = len(s)
  var indexZuo3, indexYou4 int = 0, sLen - 1
  var arr1S []byte = []byte(s)

  for indexZuo3 < indexYou4 {
    for indexZuo3 < indexYou4 && !strings.Contains("aeiouAEIOU", string(s[indexZuo3])) {
      indexZuo3++
    }
    for indexZuo3 < indexYou4 && !strings.Contains("aeiouAEIOU", string(s[indexYou4])) {
      indexYou4--
    }
    if indexZuo3 < indexYou4 {
      arr1S[indexZuo3], arr1S[indexYou4] = arr1S[indexYou4], arr1S[indexZuo3]
      indexZuo3++
      indexYou4--
    }
  }

  return string(arr1S)
}
