package main

import "fmt"

func main() {
  fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
  fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}

//给定一个非递减顺序的整数数组，返回每个数字的平方组成的新数组，要求新数组也按非递减顺序排序

//数组，排序，双指针，数学
//负数越小平方越大，正数越大平方越大
//因为数组是有序的，所以最小的负数和最大的正数分别位于两侧
//从两侧开始，将数字平方后比较，大的那个就是结果数组里大的那个

//977-有序数组的平方
func sortedSquares(nums []int) []int {
  var numsLen int = len(nums)
  var sli1Res []int = make([]int, numsLen)
  var indexRes int = numsLen - 1
  var indexSmall, indexBig int = 0, numsLen - 1

  for indexSmall <= indexBig {
    if nums[indexSmall]*nums[indexSmall] > nums[indexBig]*nums[indexBig] {
      sli1Res[indexRes] = nums[indexSmall] * nums[indexSmall]
      indexRes--
      indexSmall++
    } else {
      sli1Res[indexRes] = nums[indexBig] * nums[indexBig]
      indexRes--
      indexBig--
    }
  }

  return sli1Res
}
