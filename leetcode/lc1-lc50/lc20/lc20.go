package main

import "fmt"

func main() {
  fmt.Println(isValid("()"))
  fmt.Println(isValid("()[]{}"))
  fmt.Println(isValid("([)]"))
  fmt.Println(isValid("(]"))
  fmt.Println(isValid("{[]}"))
  fmt.Println(isValid("(])"))
}

//给定一个只包括'('，')'，'{'，'}'，'['，']'的字符串s，判断字符串是否有效。
//有效字符串需满足：左括号必须用相同类型的右括号闭合。左括号必须以正确的顺序闭合。
//1<=s.length<=10^4;s仅由括号'()[]{}'组成

//字符串，栈

//20-有效的括号
func isValid(s string) bool {
  var sLen int = len(s)
  var stack [10000]byte = [10000]byte{}
  var indexTop int = 0

  for index := 0; index < sLen; index++ {
    if s[index] == '(' || s[index] == '[' || s[index] == '{' {
      stack[indexTop] = s[index]
      indexTop++
    } else if s[index] == ')' || s[index] == ']' || s[index] == '}' {
      if indexTop <= 0 {
        return false
      }
      if s[index] == ')' && stack[indexTop-1] == '(' {
        indexTop--
      } else if s[index] == ']' && stack[indexTop-1] == '[' {
        indexTop--
      } else if s[index] == '}' && stack[indexTop-1] == '{' {
        indexTop--
      } else {
        return false
      }
    } else {
      return false
    }
  }

  return indexTop <= 0
}
