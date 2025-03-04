package main

import "fmt"

// 剑指 Offer 04. 二维数组中的查找
// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
// 请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

// 解题思路
// 从右上角开始，的每个元素，左边的元素一定比自己小，下面的元素一定比自己大。

func main() {
  matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
  fmt.Println(findNumberIn2DArray(matrix, 5))
  fmt.Println(findNumberIn2DArray(matrix, 23))
  fmt.Println(findNumberIn2DArray(matrix, 20))
  fmt.Println(findNumberIn2DArray(matrix, 25))
  fmt.Println(findNumberIn2DArray(matrix, 31))

  matrix = [][]int{{-5}}
  fmt.Println(findNumberIn2DArray(matrix, -2))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
  line := len(matrix)
  if line <= 0 {
    return false
  }
  column := len(matrix[0])
  if column <= 0 {
    return false
  }

  i := 0
  j := column - 1
  for {
    if matrix[i][j] < target {
      i++
    } else if matrix[i][j] > target {
      j--
    } else {
      return true
    }

    if i >= line || j < 0 {
      return false
    }
  }
}
