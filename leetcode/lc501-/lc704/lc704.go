package main

import "fmt"

func main() {
  fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 5))
}

//给定一个n个元素有序的（升序）整型数组nums和一个目标值target，
//写一个函数搜索nums中的target，如果目标值存在返回下标，否则返回-1。
//假设nums中的所有元素是不重复的。n将在[1,10000]之间。nums的每个元素都将在[-9999,9999]之间。

//704-二分查找
func search(nums []int, target int) int {
  indexLeft, indexRight := 0, len(nums)-1
  for indexLeft <= indexRight {
    indexMid := indexLeft + (indexRight-indexLeft)>>1 //这么写可以防止数值过大溢出
    if nums[indexMid] > target {
      indexRight = indexMid - 1 //中间数大于目标数，目标数在左半边，移动右边界
    } else if nums[indexMid] < target {
      indexLeft = indexMid + 1 //中间数小于目标数，目标数在右半边，移动左边界
    } else {
      return indexMid //找到目标
    }
  }

  return -1
}
