package main

import "fmt"

func main() {
  fmt.Println(firstUniqChar("leetcode"))
  fmt.Println(firstUniqChar("loveleetcode"))
}

//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回-1。

//字符串，哈希表
//用哈希表统计字符串中每个字符的数量
//在遍历一次字符串，在哈希表中找到的第一个数量是1的字符

//387-字符串中的第一个唯一字符
func firstUniqChar(s string) int {
  var sLen int = len(s)
  var hashChar [128]int = [128]int{}

  for index := 0; index < sLen; index++ {
    hashChar[s[index]]++
  }

  for index := 0; index < sLen; index++ {
    if hashChar[s[index]] == 1 {
      return index
    }
  }

  return -1
}
