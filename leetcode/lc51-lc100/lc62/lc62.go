package main

import "fmt"

func main() {
  fmt.Println(uniquePaths(3, 7))
}

//一个机器人位于一个m行n列的网格的左上角，机器人每次只能向下或者向右移动一步
//机器人试图达到网格的右下角，问总共有多少条不同的路径

//62-不同路径
func uniquePaths(m int, n int) int {
  //动态规划
  //到达[i,j]位置的总路径等于，到达[i,j-1]位置的总路径，加上到达[i-1,j]位置的总路径
  //第一行和第一列的格子都只有一条路径

  isli2Res := make([][]int, m)

  //初始化第一行和第一列
  for ii := 0; ii < m; ii++ {
    isli2Res[ii] = make([]int, n)
    isli2Res[ii][0] = 1
  }
  for ii := 0; ii < n; ii++ {
    isli2Res[0][ii] = 1
  }

  for ii := 1; ii < m; ii++ {
    for ij := 1; ij < n; ij++ {
      isli2Res[ii][ij] = isli2Res[ii-1][ij] + isli2Res[ii][ij-1]
    }
  }

  return isli2Res[m-1][n-1]
}
