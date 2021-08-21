package main

import (
  "fmt"
)

func main() {
  // chars1 := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
  // fmt.Println(compress(chars1))
  // fmt.Println(chars1)

  chars2 := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
  fmt.Println(compress(chars2))
  fmt.Println(chars2)
}

//给你一个字符数组chars，请使用下述算法压缩：
//从一个空字符串s开始。对于chars中的每组连续重复字符：
//如果这一组长度为1，则将字符追加到s中。否则，需要向s追加字符，后跟这一组的长度。
//压缩后得到的字符串s不应该直接返回，需要转储到字符数组chars中。
//需要注意的是，如果组长度为10或10以上，则在chars数组中会被拆分为多个字符。
//请在修改完输入数组后，返回该数组的新长度。
//你必须设计并实现一个只使用常量额外空间的算法来解决此问题。

//双指针
//一个指针用于遍历，一个指针用于重写数据
//注意数字大于9的情况，需要占多个位置

//433-压缩字符串
func compress(chars []byte) int {
  var charsLen int = len(chars)
  var indexRead, indexWrite int = 1, 0
  var charNow byte = chars[0]
  var charCount int = 1

  for ; indexRead <= charsLen; indexRead++ {
    if indexRead == charsLen || chars[indexRead] != charNow {
      if charCount == 1 {
        chars[indexWrite] = charNow
        indexWrite++
      } else {
        chars[indexWrite] = charNow
        indexWrite++
        //int转string转[]byte然后赋值
        // var charCountstr string = strconv.FormatInt(int64(charCount), 10)
        // var arrCharCountstr []byte = []byte(charCountstr)
        // for index := 0; index < len(arrCharCountstr); index++ {
        //   chars[indexWrite] = arrCharCountstr[index]
        //   indexWrite++
        // }
        //当然也可以对10取余反着算
        var charCountStr [10]byte = [10]byte{}
        var charCountStrLen int = 0
        for ; charCount > 0; charCount /= 10 {
          charCountStr[charCountStrLen] = '0' + byte(charCount%10)
          charCountStrLen++
        }
        for ; charCountStrLen > 0; charCountStrLen-- {
          chars[indexWrite] = charCountStr[charCountStrLen-1]
          indexWrite++
        }
      }
      if indexRead < charsLen {
        charNow = chars[indexRead]
        charCount = 1
      }
      continue
    }
    charCount++
  }

  return indexWrite
}
