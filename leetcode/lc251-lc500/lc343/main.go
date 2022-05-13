package main

import "fmt"

// 343. 整数拆分
// 给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
// 返回 你可以获得的最大乘积 。

// 解题思路：数学
// n=n1+n1+n3+...+na，求 max(n1*n2*n3*...na)
// 算数几何不等式：(n1+n1+n3+...+na)/a >= a√(n1*n2*n3*...na)
// 如果将绳子等分为 a 段，每段长为 x。则最后的结论是：x=e≈2.7 时，沉积最大。
// 3 更接近 e，所以整数长度取 3 而不是 2。最后注意一下 4，2*2>1*3。

func main() {
  fmt.Println(integerBreak(2))
  fmt.Println(integerBreak(3))
  fmt.Println(integerBreak(4))
  fmt.Println(integerBreak(5))
  fmt.Println(integerBreak(7))
  fmt.Println(integerBreak(10))
}

func integerBreak(n int) int {
  // 前面两个需要特别处理
  if 2 == n {
    return 1
  } else if 3 == n {
    return 2
  }
  sum := 1
  a := n
  for {
    if a > 4 {
      sum *= 3
      a -= 3
    } else {
      sum *= a
      break
    }
  }
  return sum
}
