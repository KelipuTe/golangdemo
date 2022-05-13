package main

import "fmt"

// 剑指 Offer 14- II. 剪绳子 II
// 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。
// 请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

// 解题思路
// 注意这里数字会溢出，需要求余，(x*y)%p = [(x%p)*(y%p)]%p

func main() {
  fmt.Println(cuttingRope(120))
}

func cuttingRope(n int) int {
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
      sum = (sum % 1000000007 * 3) % 1000000007
      a -= 3
    } else {
      sum = (sum % 1000000007 * a) % 1000000007
      break
    }
  }
  return sum
}
