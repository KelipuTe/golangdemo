package main

import "fmt"

// 70. 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

// 爬到 n 阶楼梯有两种办法，从 n-1 阶楼梯爬 1 阶，或者，从 n-2 阶楼梯爬 2 阶
// 所以，爬到 n 阶楼梯的结果，等于爬到 n-1 阶楼梯的结果，和爬到 n-2 阶楼梯的结果之和

func main() {
  fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
  if 0 == n {
    return 0
  } else if 1 == n {
    return 1
  } else if 2 == n {
    return 2
  }

  sli1num := make([]int, n+1)
  sli1num[0] = 0
  sli1num[1] = 1
  sli1num[2] = 2
  for i := 3; i <= n; i++ {
    sli1num[i] = sli1num[i-1] + sli1num[i-2]
  }

  return sli1num[n]
}
