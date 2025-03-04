package main

import (
  "fmt"
)

func main() {
  fmt.Println(isAnagram("anagram", "nagaram"))
  fmt.Println(isAnagram("rat", "car"))
}

//若s和t中每个字符出现的次数都相同，则称s和t互为字母异位词
//给定两个字符串s和t，判断t是否是s的字母异位词
//1<=s.length, t.length<=5*10^4;s和t仅包含小写字母

//字符串，哈希表
//遍历字符串s，用哈希表统计每个字符出现的次数
//遍历字符串t，遇到一个字符，就在哈希表中对应的位置-1
//如果哈希表所有位置都是0，则s和t是字母异位词

//242-有效的字母异位词
func isAnagram(s string, t string) bool {
  var sLen, tLen int = len(s), len(t)
  var hashChar [128]int = [128]int{}

  for index := 0; index < sLen; index++ {
    hashChar[s[index]]++
  }

  for index := 0; index < tLen; index++ {
    hashChar[t[index]]--
  }

  checkTemp := true
  //题目中只有小写字母，所以检查范围是，从a=97到z=122
  for index := 97; index < 123; index++ {
    if hashChar[index] != 0 {
      checkTemp = false
      break
    }
  }

  return checkTemp
}
