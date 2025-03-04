package main

import "fmt"

func main() {
  fmt.Println(checkInclusion("ab", "eidbaooo"))
  fmt.Println(checkInclusion("ab", "eidboaoo"))
  fmt.Println(checkInclusion("aoo", "eidboaoo"))
  fmt.Println(checkInclusion("eib", "eidboaoo"))
  fmt.Println(checkInclusion("adc", "dcda"))
  fmt.Println(checkInclusion("abc", "bbbca"))
}

//给你两个字符串s1和s2，写一个函数来判断s2是否包含s1的排列。
//换句话说，s1的排列之一是s2的子串。
//1<=s1.length,s2.length<=10^4;s1和s2仅包含小写字母;

//字符串，哈希表，滑动窗口
//遍历s1，构造哈希表，统计s1中出现的字符串
//使用左右两个指针构造一个滑动窗口，处理s2。右指针先遍历，直到窗口长度等于s1长度。
//判断窗口中字符是否匹配，这里需要将上面的哈希表复制一份。
//如果窗口中的字符属于s1，哈希表对应位置就-1，如果哈希表全部为0证明匹配成功。
//将窗口向右整体移动一个单位长度，直到匹配成功或者窗口移动到s2末尾。

//567-字符串的排列
func checkInclusion(s1 string, s2 string) bool {
  var s1Len, s2Len int = len(s1), len(s2)
  var hashCharS1 [128]int = [128]int{}   //记录s1字符数量
  var hashCharTemp [128]int = [128]int{} //用于判断时，动态标记剩余待匹配字符
  var indexi, indexj int = 0, 0          //滑动窗口
  var lenTemp int = 0                    //滑动窗口长度

  if s1Len > s2Len {
    return false
  }

  for index := 0; index < s1Len; index++ {
    hashCharS1[s1[index]]++
    hashCharTemp[s1[index]]++
  }

  for indexj < s2Len {
    //先把窗口长度增加到s1字符串长度
    for lenTemp < s1Len-1 {
      if hashCharS1[s2[indexj]] > 0 {
        //匹配到字符，哈希表对应位置-1
        hashCharTemp[s2[indexj]]--
      }
      lenTemp++ //没匹配到字符，窗口长度+1
      indexj++
    }
    //这时窗口长度为s1字符串长度-1
    if hashCharS1[s2[indexj]] > 0 {
      hashCharTemp[s2[indexj]]--
    }
    indexj++

    checkTemp := true
    //题目中只有小写字母，所以检查范围是，从a=97到z=122
    for index := 97; index < 123; index++ {
      if hashCharTemp[index] != 0 {
        //全0才算匹配，这里有可能会减成负数
        checkTemp = false
        break
      }
    }
    if checkTemp {
      return true
    }

    if hashCharTemp[s2[indexi]] < hashCharS1[s2[indexi]] {
      //如果左指针指向的字符是s1中的字符，则哈希表对应位置+1
      hashCharTemp[s2[indexi]]++
    }
    indexi++
  }

  return false
}
