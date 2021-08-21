package main

import (
  "fmt"
)

func main() {
  fmt.Println(searchMatrix([][]int{
    {1, 4, 7, 11, 15},
    {2, 5, 8, 12, 19},
    {3, 6, 9, 16, 22},
    {10, 13, 14, 17, 24},
    {18, 21, 23, 26, 30},
  }, 8))

  fmt.Println(searchMatrix([][]int{
    {1, 4, 7, 11, 15},
    {2, 5, 8, 12, 19},
    {3, 6, 9, 16, 22},
    {10, 13, 14, 17, 24},
    {18, 21, 23, 26, 30},
  }, 20))
}

//编写一个高效的算法来搜索mxn矩阵matrix中的一个目标值target。该矩阵具有以下特性：
//每行的元素从左到右升序排列。每列的元素从上到下升序排列。
//m==matrix.length;n==matrix[i].length;1<=n,m<=300
//-10^9<=matix[i][j]<=10^9;-10^9<=target<=10^9
//每行的所有元素从左到右升序排列;每列的所有元素从上到下升序排列

//根据矩阵的特性，可以知道，左下角那个元素，比它在的那一列都大，比它在的那一行都小
//所以，每次判断都可以缩小查询范围，缩小一行或者一列的，然后重新选定新的左下角

//240-搜索二维矩阵II
func searchMatrix(matrix [][]int, target int) bool {
  var hang2shu4, lie4shu4 int = len(matrix), len(matrix[0])

  for m, n := hang2shu4-1, 0; m >= 0 && n < lie4shu4; {
    var zuo3xia4num int = matrix[m][n] //左下角元素
    if zuo3xia4num > target {
      m--
    } else if zuo3xia4num < target {
      n++
    } else {
      return true //找到
    }
  }

  return false
}
