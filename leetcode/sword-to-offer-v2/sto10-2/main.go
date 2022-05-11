package main

import (
  "fmt"
)

// 剑指 Offer 10- II. 青蛙跳台阶问题
// 一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

func main() {
  fmt.Println(numWays(0))
  fmt.Println(numWays(1))
  fmt.Println(numWays(2))
  fmt.Println(numWays(7))
  fmt.Println(numWays(100))
}

func numWays(n int) int {
  if 0 == n {
    return 1
  } else if 1 == n {
    return 1
  } else if 2 == n {
    return 2
  }

  sli1num := make([]int64, n+1)
  sli1num[0] = 1
  sli1num[1] = 1
  sli1num[2] = 2
  for i := 3; i <= n; i++ {
    sli1num[i] = sli1num[i-1]%1000000007 + sli1num[i-2]%1000000007
  }

  return int(sli1num[n] % 1000000007)
}
