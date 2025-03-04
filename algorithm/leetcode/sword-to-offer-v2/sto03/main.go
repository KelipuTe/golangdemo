package main

import "fmt"

// 剑指 Offer 03. 数组中重复的数字

// 解题思路
// 哈希表省时间，先排序再判断省空间

func main() {
  nums := []int{3, 4, 2, 0, 0, 1}
  fmt.Println(findRepeatNumber(nums))
}

func findRepeatNumber(nums []int) int {
  numsLen := len(nums)
  mapexist := make(map[int]bool, numsLen)
  for i := 0; i < numsLen; i++ {
    if _, ok := mapexist[nums[i]]; ok {
      return nums[i]
    } else {
      mapexist[nums[i]] = true
    }
  }
  return -1
}
