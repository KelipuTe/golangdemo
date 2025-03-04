package main

import "fmt"

func main() {
  fmt.Println(compareVersion("1.01", "1.001"))
  fmt.Println(compareVersion("1.0", "1.0.0"))
  fmt.Println(compareVersion("0.1", "1.1"))
  fmt.Println(compareVersion("1.0.1", "1"))
  fmt.Println(compareVersion("7.5.2.4", "7.5.3"))
}

//给你两个版本号version1和version2，请你比较它们。
//版本号由一个或多个修订号组成，各修订号由一个'.'连接。每个修订号由多位数字组成，可能包含前导零。每个版本号至少包含一个字符。
//修订号从左到右编号，下标从0开始，最左边的修订号下标为0，下一个修订号下标为1，以此类推。例如，2.5.33和0.1都是有效的版本号。
//比较版本号时，请按从左到右的顺序依次比较它们的修订号。比较修订号时，只需比较忽略任何前导零后的整数值。
//也就是说，修订号1和修订号001相等。如果版本号没有指定某个下标处的修订号，则该修订号视为0。
//例如，版本1.0小于版本1.1，因为它们下标为0的修订号相同，而下标为1的修订号分别为0和1，0<1。
//返回规则：如果version1>version2返回1，如果version1<version2返回-1，除此之外返回0。

//双指针
//两个指针分别遍历字符串，把版本号算成数字，然后比较这两个数字的大小
//两个版本不一样长的时候，缺失的那个版本号默认为0，这样方便写逻辑

//165-比较版本号
func compareVersion(version1 string, version2 string) int {
  var version1Len, version2Len int = len(version1), len(version2)

  index1, index2 := 0, 0
  for index1 < version1Len || index2 < version2Len {
    num1 := 0
    for ; index1 < version1Len && version1[index1] != '.'; index1++ {
      num1 *= 10
      num1 += (int)(version1[index1] - '0')
    }
    index1++

    num2 := 0
    for ; index2 < version2Len && version2[index2] != '.'; index2++ {
      num2 *= 10
      num2 += (int)(version2[index2] - '0')
    }
    index2++

    if num1 > num2 {
      return 1
    } else if num1 < num2 {
      return -1
    }
  }

  return 0
}
