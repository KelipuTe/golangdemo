package main

import "fmt"

func main() {
  fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

//给定一个包含非负整数的m行n列网格grid ，每次只能向下或者向右移动一步
//请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小

//64-最小路径和(62,64)
func minPathSum(grid [][]int) int {
  //动态规划
  //和第62题的思路类似，要计算到达[i,j]位置的最小路径和，依赖于
  //到达[i,j-1]位置的最小路径和，和到达[i-1,j]位置的最小路径和
  //这两个结果中小的那个，再加上[i,j]位置的数值
  //第一行和第一列的格子都只有一条路径，所以依次累加就能得到结果

  m, n := len(grid), len(grid[0])
  isli2Res := make([][]int, m)

  //初始化第一行和第一列
  for ii := 0; ii < m; ii++ {
    if ii == 0 {
      isli2Res[ii] = make([]int, n)
      isli2Res[ii][0] = grid[ii][0]
    } else {
      isli2Res[ii] = make([]int, n)
      isli2Res[ii][0] = isli2Res[ii-1][0] + grid[ii][0]
    }
  }
  for ii := 0; ii < n; ii++ {
    if ii == 0 {
      isli2Res[0][ii] = grid[0][ii]
    } else {
      isli2Res[0][ii] = isli2Res[0][ii-1] + grid[0][ii]
    }
  }

  for ii := 1; ii < m; ii++ {
    for ij := 1; ij < n; ij++ {
      if isli2Res[ii-1][ij] < isli2Res[ii][ij-1] {
        isli2Res[ii][ij] = isli2Res[ii-1][ij] + grid[ii][ij]
      } else {
        isli2Res[ii][ij] = isli2Res[ii][ij-1] + grid[ii][ij]
      }
    }
  }

  return isli2Res[m-1][n-1]
}
