package main

import (
  "fmt"
)

func main() {
  fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
  // fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}

//给你一个整数数组nums，找到其中最长严格递增子序列的长度。
//子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
//例如[3,6,2,7]是数组[0,3,1,6,2,2,7]的子序列。

//动态规划
//构造一个数组，保存数组从下标0到下标n为止最长的子序列的长度。
//下标n的结果，等于前面n-1个下标中，满足条件nums[n-1]<nums[n]的最长的子序列的长度+1。
//两个判断条件，在构造的数组中，要找最长的子序列，在原数组中，对应位置的数要比当前位置的数小。

//300-最长递增子序列
func lengthOfLIS(nums []int) int {
  var numsLen = len(nums)
  var sli1Res []int
  var maxLen int = 1

  sli1Res = make([]int, numsLen)
  for indexNow := 0; indexNow < numsLen; indexNow++ {
    sli1Res[indexNow] = 1 //初始化1
    for indexFront := 0; indexFront < indexNow; indexFront++ {
      if nums[indexFront] < nums[indexNow] && sli1Res[indexFront]+1 > sli1Res[indexNow] {
        sli1Res[indexNow] = sli1Res[indexFront] + 1
        if sli1Res[indexNow] > maxLen {
          maxLen = sli1Res[indexNow]
        }
      }
    }
  }

  return maxLen
}
