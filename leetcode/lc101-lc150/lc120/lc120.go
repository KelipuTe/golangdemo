//lc120-三角形最小路径和
//[动态规划]

//给定一个三角形triangle，找出自顶向下的最小路径和。
//每一步只能移动到下一行中相邻的结点上。
//相邻的结点在这里指的是下标与上一层结点下标相同或者等于上一层结点下标+1的两个结点。
//也就是说，如果正位于当前行的下标i，那么下一步可以移动到下一行的下标i或i+1。
//进阶：只使用O(n)的额外空间（n为三角形的总行数）来解决这个问题。

//下一行的路径最小和的结果取决于上一行的结果。
//下一行的第一个位置和最后一个位置，只依赖上一行的第一个位置和最后一个位置的结果。
//下一行中间的位置，可以从上一行的两个位置过来，需要计算出小的那个。
//进阶，只用一维数组时，每行从后往前算，可以解决数据被覆盖的问题。

package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
  fmt.Println(minimumTotal([][]int{{-10}}))
}

func minimumTotal(triangle [][]int) int {
  var hang2 int = len(triangle)
  var sli1Res []int = make([]int, hang2)
  var pathMin int = math.MaxInt32 //初始化一个极大值

  //计算最后一行的结果
  sli1Res[0] = triangle[0][0]
  for hang2Temp := 1; hang2Temp < hang2; hang2Temp++ {
    sli1Res[hang2Temp] = sli1Res[hang2Temp-1] + triangle[hang2Temp][hang2Temp] //第一个
    for lie4Temp := hang2Temp - 1; lie4Temp > 0; lie4Temp-- {                  //中间的倒着算
      num1 := sli1Res[lie4Temp] + triangle[hang2Temp][lie4Temp]
      num2 := sli1Res[lie4Temp-1] + triangle[hang2Temp][lie4Temp]
      if num1 < num2 {
        sli1Res[lie4Temp] = num1
      } else {
        sli1Res[lie4Temp] = num2
      }
    }
    sli1Res[0] = sli1Res[0] + triangle[hang2Temp][0] //最后一个
  }
  //找最小的
  for index := 0; index < hang2; index++ {
    if sli1Res[index] < pathMin {
      pathMin = sli1Res[index]
    }
  }

  return pathMin
}
