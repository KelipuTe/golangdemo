//lc70-爬楼梯
//[动态规划]

//假设正在爬楼梯。需要n阶才能到达楼顶。
//每次可以爬1或2个台阶。有多少种不同的方法可以爬到楼顶呢？
//注意：给定n是一个正整数。

//爬到n阶楼梯有两种办法，从n-1阶楼梯爬1阶，或者，从n-2阶楼梯爬2阶
//所以，爬到n阶楼梯的结果，等于爬到n-1阶楼梯的结果，和爬到n-2阶楼梯的结果之和

package main

import "fmt"

func main() {
  fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
  var sli1Res []int = make([]int, n+1)

  if n == 1 {
    return 1
  }
  if n == 2 {
    return 2
  }

  sli1Res[0] = 0
  sli1Res[1] = 1
  sli1Res[2] = 2
  for index := 3; index <= n; index++ {
    sli1Res[index] = sli1Res[index-1] + sli1Res[index-2]
  }

  return sli1Res[n]
}
