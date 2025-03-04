package main

import "fmt"

func main() {
  fmt.Println(runningSum([]int{2, 7, 11, 15}))
}

//数组动态和的计算公式为：runningSum[i]=sum(nums[0]…nums[i])。
//给一个数组nums。请返回nums的动态和。
//1<=nums.length<=1000;-10^6<=nums[i]<=10^6

//1480-一维数组的动态和
func runningSum(nums []int) []int {
  var numsLen = len(nums)
  var sli1Res []int = make([]int, numsLen)
  var sumRes int

  for index := 0; index < numsLen; index++ {
    sumRes += nums[index]
    sli1Res[index] = sumRes
  }

  return sli1Res
}
