package main

import "fmt"

type ListNode struct {
  Val  int
  Next *ListNode
}

func main() {
  var arr2 [][]int

  arr2 = [][]int{
    {1, 1, 1},
    {1, 0, 1},
    {1, 1, 1},
  }
  setZeroes(arr2)
  fmt.Println(arr2)

  arr2 = [][]int{
    {0, 1, 2, 0},
    {3, 4, 5, 2},
    {1, 3, 1, 5},
  }
  setZeroes(arr2)
  fmt.Println(arr2)
}

//给定一个mxn的矩阵，如果一个元素为0，则将其所在行和列的所有元素都设为0。
//请使用原地算法。不要重新申请空间保存矩阵。
//进阶：仅使用常量空间的解决方案。
//m==matrix.length;n==matrix[0].length;
//1<=m,n<=200;-2^31<=matrix[i][j]<=2^31-1

func setZeroes(matrix [][]int) {
  var hang2 int = len(matrix)
  var lie4 int = len(matrix[0])
  var sli1hang2 []bool = make([]bool, hang2)
  var sli1lie4 []bool = make([]bool, lie4)

  //判断哪行那列有0
  for indexi := 0; indexi < hang2; indexi++ {
    for indexj := 0; indexj < lie4; indexj++ {
      if matrix[indexi][indexj] == 0 {
        sli1hang2[indexi] = true
        sli1lie4[indexj] = true
      }
    }
  }
  //重新赋值
  for indexi := 0; indexi < hang2; indexi++ {
    for indexj := 0; indexj < lie4; indexj++ {
      if sli1hang2[indexi] || sli1lie4[indexj] {
        matrix[indexi][indexj] = 0
      }
    }
  }
}
