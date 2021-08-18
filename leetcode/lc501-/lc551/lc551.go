package main

import "fmt"

func main() {
  fmt.Println(checkRecord("PPALLP"))
  fmt.Println(checkRecord("PPALLL"))
}

//给你一个字符串s表示一个学生的出勤记录，其中的每个字符用来标记当天的出勤情况（缺勤、迟到、到场）。记录中只含下面三种字符：
//'A'：Absent，缺勤；'L'：Late，迟到；'P'：Present，到场；
//如果学生能够同时满足下面两个条件，则可以获得出勤奖励：
//按总出勤计，学生缺勤（'A'）严格少于两天。学生不会存在连续3天或3天以上的迟到（'L'）记录。
//如果学生可以获得出勤奖励，返回true；否则，返回false。
//1<=s.length<=1000;[i]为'A'、'L'或'P';

//551-学生出勤记录I
func checkRecord(s string) bool {
  iSumA, iLineL, iMaxLineL := 0, 0, 0

  for iIndex := 0; iIndex < len(s); iIndex++ {
    switch s[iIndex] {
    case 'A':
      iSumA++
      if iMaxLineL < iLineL {
        iMaxLineL = iLineL
      }
      iLineL = 0
      break
    case 'L':
      iLineL++
      if iMaxLineL < iLineL {
        iMaxLineL = iLineL
      }
      break
    case 'P':
      if iMaxLineL < iLineL {
        iMaxLineL = iLineL
      }
      iLineL = 0
      break
    }
  }

  if iSumA < 2 && iMaxLineL <= 2 {
    return true
  }
  return false
}
