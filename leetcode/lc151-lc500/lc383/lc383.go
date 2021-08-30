package main

import (
  "fmt"
)

func main() {
  fmt.Println(canConstruct("a", "b"))
  fmt.Println(canConstruct("aa", "ab"))
  fmt.Println(canConstruct("aa", "aab"))
}

//给定一个赎金信(ransom)字符串和一个杂志(magazine)字符串，
//为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。
//杂志字符串中的每个字符只能在赎金信字符串中使用一次。
//判断第一个字符串ransom能不能由第二个字符串magazines里面的字符构成。
//如果可以构成，返回true；否则返回false。
//可以假设两个字符串均只含有小写字母。

//字符串，哈希表
//遍历杂志，用哈希表统计每个字符出现的次数
//遍历赎金信，遇到一个字符，就在哈希表中对应的位置-1
//如果哈希表所有位置都大于0，则可以构成

//383-赎金信
func canConstruct(ransomNote string, magazine string) bool {
  var ransomNoteLen, magazineLen int = len(ransomNote), len(magazine)
  var hashChar [128]int = [128]int{}

  for index := 0; index < magazineLen; index++ {
    hashChar[magazine[index]]++
  }

  for index := 0; index < ransomNoteLen; index++ {
    hashChar[ransomNote[index]]--
  }

  checkTemp := true
  //题目中只有小写字母，所以检查范围是，从a=97到z=122
  for index := 97; index < 123; index++ {
    if hashChar[index] < 0 {
      checkTemp = false
      break
    }
  }

  return checkTemp
}
