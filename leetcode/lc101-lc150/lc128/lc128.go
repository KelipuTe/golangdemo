package main

import "fmt"

func main() {
  fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
  fmt.Println(longestConsecutive([]int{1, 2}))
  fmt.Println(longestConsecutive([]int{1}))
  fmt.Println(longestConsecutive([]int{1, 1}))
}

//给定一个未排序的整数数组nums，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度
//请设计并实现时间复杂度为O(n)的算法解决此问题

//哈希表
//这里要求时间复杂度为O(n)，所以先排序后遍历计算是不行的
//将数组转换成哈希表之后，可以方便查询某一个数字前后的数字是否存在
//基本思路是遍历每一个数字，然后向后查询数字是否存在，但是这样的时间复杂度为O(n^2)
//在遍历的过程中可以发现，如果存在，x,x+1,x+2，那么从x+1开始的序列其实是不需要遍历的
//所以可以规定，如果一个数字x，前面一个数字x-1，不存在，则从x-1开始查询，否则不查询
//这样，x,x+1,x+2，这个序列，只会在遍历到x时遍历一次

//128-最长连续序列
func longestConsecutive(nums []int) int {
  iNumsLen := len(nums)
  if iNumsLen < 1 {
    return 0
  }

  var mapNums map[int]bool = map[int]bool{}
  for ii := 0; ii < iNumsLen; ii++ {
    mapNums[nums[ii]] = true
  }

  iLongestRes := 1
  for iNum := range mapNums {
    if mapNums[iNum-1] {
      continue
    }
    iCurrentRes := 1
    for mapNums[iNum+1] {
      iNum++
      iCurrentRes++
    }
    if iCurrentRes > iLongestRes {
      iLongestRes = iCurrentRes
    }
  }

  return iLongestRes
}
