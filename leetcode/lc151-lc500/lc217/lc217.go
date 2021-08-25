package main

import "fmt"

func main() {
  fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
  fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

//给定一个整数数组，判断是否存在重复元素。
//如果存在一值在数组中出现至少两次，函数返回true。
//如果数组中每个元素都不相同，则返回false。

//哈希表
//遍历数组并用哈希表计数，遍历哈希表，判断有没有计数大于1的

//217-存在重复元素
func containsDuplicate(nums []int) bool {
  var mapNums map[int]int = map[int]int{}

  for index := 0; index < len(nums); index++ {
    mapNums[nums[index]]++
  }

  for _, value := range mapNums {
    if value > 1 {
      return true
    }
  }

  return false
}
