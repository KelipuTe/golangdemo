package main

import "fmt"

func main() {
  fmt.Println(generate(1))
  fmt.Println(generate(2))
  fmt.Println(generate(3))
  fmt.Println(generate(4))
  fmt.Println(generate(5))
}

//给定一个非负整数numRows，生成杨辉三角的前numRows行。
//在杨辉三角中，每个数是它左上方和右上方的数的和。

//118-杨辉三角
func generate(numRows int) [][]int {
  var sli2Res [][]int = make([][]int, numRows)

  for line := 1; line <= numRows; line++ {
    sli2Res[line-1] = make([]int, line)
    for index := 0; index < line; index++ {
      if index == 0 || index == line-1 {
        sli2Res[line-1][index] = 1
      } else {
        sli2Res[line-1][index] = sli2Res[line-2][index-1] + sli2Res[line-2][index]
      }
    }
  }

  return sli2Res
}
