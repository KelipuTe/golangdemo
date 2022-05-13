package main

import "fmt"

// 剑指 Offer 13. 机器人的运动范围
// 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
// 一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
// 例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

// 解题思路：动态规划
// 可以用下面的 printMap 函数，输出一下地图，地图是有规律的。每一行和每一列都是阶段性递增的关系。
// 如果一个点能到达，那么一定是从左边或者上边过来的；能从右边或者下面过来的，也一定能从左边或者上边过来。
// 所以问题可以简化成：每一个格子能不能从左边或者上面过来。

func main() {
  // fmt.Println(movingCount(2, 3, 1))
  // fmt.Println(movingCount(3, 1, 0))
  // fmt.Println(movingCount(16, 8, 4))
  fmt.Println(movingCount(14, 14, 5))
}

func movingCount(m int, n int, k int) int {
  sli2isVisit := make([][]bool, m)
  for i := 0; i < m; i++ {
    sli2isVisit[i] = make([]bool, n)
  }
  visitNum := 0
  for i := 0; i < m; i++ {
    for j := 0; j < n; j++ {
      if 0 == i && 0 == j {
        sli2isVisit[i][j] = true
        visitNum++
        continue
      }
      // 计算下标之和
      sum := 0
      a := i
      for a > 0 {
        sum += a % 10
        a = a / 10
      }
      b := j
      for b > 0 {
        sum += b % 10
        b = b / 10
      }
      if k >= sum {
        // 第一行
        if 0 == i {
          if sli2isVisit[0][j-1] {
            sli2isVisit[i][j] = true
            visitNum++
          }
          continue
        }
        // 第一列
        if 0 == j {
          if sli2isVisit[i-1][0] {
            sli2isVisit[i][j] = true
            visitNum++
          }
          continue
        }
        // 中间的
        if sli2isVisit[i-1][j] || sli2isVisit[i][j-1] {
          sli2isVisit[i][j] = true
          visitNum++
          continue
        }
      }
    }
  }
  return visitNum
}

// 输出地图，地图其实是有规律的
func printMap() {
  for i := 0; i < 20; i++ {
    for j := 0; j < 20; j++ {
      sum := 0
      m := i
      n := j
      for m > 0 {
        sum += m % 10
        m = m / 10
      }
      for n > 0 {
        sum += n % 10
        n = n / 10
      }
      fmt.Printf("%3d,", sum)
    }
    fmt.Print("\r\n")
  }
}
