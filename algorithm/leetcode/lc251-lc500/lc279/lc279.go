package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println(numSquares(9))
}

//完全平方数是一个整数，其值等于另一个整数的平方。
//给定正整数n，找到若干个完全平方数（比如1,4,9,16,...）使得它们的和等于n。
//需要让组成和的完全平方数的个数最少。
//1<=n<=10^4

//动态规划
//举例讨论5的问题。1=[1],2=[1,1],3=[1,1,1],4=[4],5=[4,1]
//因为完全平方数的个数要最少，所以要穷举5之前每个完全平方数，找到个数最小的组合。
//从1开始，5的问题变成变成1+4的问题，问题没变，规模变小了。
//4的问题变成1+3的问题，3的问题变成1+2的问题，2的问题变成1+1的问题，最优解为5个数。
//注意，4的问题还可以变成4+0的问题，所以从1开始的最优解为2个数。
//从4开始，5的问题变成变成4+1的问题，和上面类似，最优解为2个数。
//解决问题的过程存在向前的依赖关系，也就是说n的问题，
//等于[1+(n-1)的问题，4+(n-4)的问题，9+(n-9)的问题...]这些问题中的最优解。
//这种问题结构可以用动态规划正向求解解决。

//279-完全平方数
func numSquares(n int) int {
  var sli1Res []int = make([]int, n+1)
  var minCount int

  sli1Res[0] = 0
  for targetNum := 1; targetNum <= n; targetNum++ {
    minCount = math.MaxInt32 //默认一个临界值
    for i := 1; i*i <= targetNum; i++ {
      minCountTemp := 1 + sli1Res[targetNum-i*i]
      if minCountTemp < minCount {
        minCount = minCountTemp
      }
    }
    sli1Res[targetNum] = minCount
  }

  return sli1Res[n]
}
