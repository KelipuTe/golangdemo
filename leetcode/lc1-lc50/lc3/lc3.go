package main

import "fmt"

func main() {
  fmt.Println(lengthOfLongestSubstring("abcabcbb"))
  fmt.Println(lengthOfLongestSubstring("bbbbb"))
  fmt.Println(lengthOfLongestSubstring("pwwkew"))
  fmt.Println(lengthOfLongestSubstring(""))
  fmt.Println(lengthOfLongestSubstring("dvdf"))
}

//给定一个字符串s，请你找出其中不含有重复字符的最长子串的长度。

//字符串，哈希表，滑动窗口
//使用左右两个指针构造一个滑动窗口，使用一个哈希表记录出现过的字符。
//右指针先遍历，直到遇到出现过的字符，这时候，右指针指向的字符，就是窗口中存在的重复的字符。
//然后，左指针开始遍历，直到遇到当前右指针指向的字符，跳过这个字符。
//这时，窗口中就没有重复的字符了，右指针继续遍历。

//3-无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
  var sLen int = len(s)
  var hashchar [128]int = [128]int{} //哈希表
  var indexi, indexj int = 0, 0 //滑动窗口
  var lenMax, lenTemp int = 0, 0

  for indexj < sLen {
    if hashchar[s[indexj]] != 1 {
      hashchar[s[indexj]] = 1
      lenTemp++
      if lenTemp > lenMax {
        lenMax = lenTemp
      }
      indexj++
    } else {
      for hashchar[s[indexi]] != hashchar[s[indexj]] {
        hashchar[s[indexi]] = 0
        lenTemp--
        indexi++
      }
      hashchar[s[indexi]] = 0
      lenTemp--
      indexi++
    }
  }

  return lenMax
}
