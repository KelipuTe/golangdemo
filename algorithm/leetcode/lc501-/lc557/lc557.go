package main

import "fmt"

func main() {
  fmt.Println(reverseWords("Let's take LeetCode contest"))
  fmt.Println(reverseWords("Let'stakeLeetCodecontest"))
  fmt.Println(reverseWords("Let'stakeLeet    Codecontest"))
}

//给定一个字符串，需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

//找到每一个单词的起始和结束位置，然后反转即可
//注意没有空格，多个空格，和末尾单词的处理

//557-反转字符串中的单词III
func reverseWords(s string) string {
  var sLen = len(s)
  var sli1Chars []byte = []byte(s)
  var indexStart, indexEnd = 0, 0

  for index := 0; index <= sLen; index++ {
    if index == sLen || sli1Chars[index] == ' ' {
      indexEnd = index - 1
      for indexTemp := indexStart; indexTemp < indexEnd-(indexEnd-indexStart)>>1; indexTemp++ {
        sli1Chars[indexTemp], sli1Chars[indexEnd-(indexTemp-indexStart)] = sli1Chars[indexEnd-(indexTemp-indexStart)], sli1Chars[indexTemp]
      }
      //跳过空格
      if index == sLen {
        break
      }
      for sli1Chars[index] == ' ' {
        index++
      }
      indexStart = index
    }
  }

  return string(sli1Chars)
}
