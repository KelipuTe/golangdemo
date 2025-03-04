package main

import (
  "fmt"
)

// 剑指 Offer 10- I. 斐波那契数列
// 写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

func main() {
  fmt.Println(fib(0))
  fmt.Println(fib(1))
  fmt.Println(fib(2))
  fmt.Println(fib(5))
  fmt.Println(fib(100))
}

func fib(n int) int {
  if 0 == n {
    return 0
  } else if 1 == n {
    return 1
  }
  sli1num := make([]int64, n+1)
  sli1num[0] = 0
  sli1num[1] = 1
  for i := 2; i <= n; i++ {
    // 这里要处理一下，要不然会溢出
    sli1num[i] = sli1num[i-2]%1000000007 + sli1num[i-1]%1000000007
  }
  return int(sli1num[n] % 1000000007)
}
