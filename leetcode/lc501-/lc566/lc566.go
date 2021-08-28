package main

import "fmt"

func main() {
  fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
  fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 2))
  fmt.Println(matrixReshape([][]int{{1, 2, 3}, {4, 5, 6}}, 3, 2))
  fmt.Println(matrixReshape([][]int{{1, 2, 3}, {4, 5, 6}}, 2, 4))
}

//在MATLAB中，有一个非常有用的函数reshape，
//它可以将一个mxn矩阵重塑为另一个大小不同（rxc）的新矩阵，但保留其原始数据。
//给你一个由二维数组mat表示的mxn矩阵，以及两个正整数r和c，分别表示想要的重构的矩阵的行数和列数。
//重构后的矩阵需要将原始矩阵的所有元素以相同的行遍历顺序填充。
//如果具有给定参数的reshape操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。

//数组，矩阵
//将原m行n列的二维数组先转换成一维数组，然后再用一维数组去填充新的二维数组。
//原二维数组第i行第j列的元素，在一维数组中就是第(i-1)*n+j个元素。
//假设一维数组中第i个元素，用i除以列数，得到的商+1就是第几行，余数就是第几列。
//注意转换成数组下标的时候是从0开始的。

//566-重塑矩阵
func matrixReshape(nums [][]int, r int, c int) [][]int {
  var hang2, lie4 int = len(nums), len(nums[0])
  var sli2Res [][]int

  if hang2*lie4 != r*c {
    return nums
  }

  sli2Res = make([][]int, r)
  for index := 0; index < r; index++ {
    sli2Res[index] = make([]int, c)
  }

  for index := 0; index < hang2*lie4; index++ {
    sli2Res[index/c][index%c] = nums[index/lie4][index%lie4]
  }

  return sli2Res
}
