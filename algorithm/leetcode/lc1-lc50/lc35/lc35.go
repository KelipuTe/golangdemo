package main

import "fmt"

func main() {
  fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
  fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
  fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
  fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
  fmt.Println(searchInsert([]int{1}, 0))
}

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
//如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

//二分查找
//找到小于目标数的最大数，然后判断这个最大数和它右边的数，哪个符合条件。
//注意目标位于边界的时候的判断。

//35-搜索插入位置
func searchInsert(nums []int, target int) int {
  numsLen := len(nums)
  index := -1 //小于目标数的最大数的下标

  indexLeft, indexRight := 0, numsLen-1
  for indexLeft <= indexRight {
    indexMid := indexLeft + (indexRight-indexLeft)>>1
    if nums[indexMid] < target {
      indexLeft = indexMid + 1 //目标数大于中间数，目标数在右半边，收缩左边界
      index = indexMid         //需要收缩左边界，则说明小于目标数的最大的元素在中间数位置或者中间数右侧位置
    } else {
      indexRight = indexMid - 1 //目标数小于中间数，目标数在左半边，收缩右边界
    }
  }

  if index == -1 {
    return 0
  }
  if index+1 > numsLen-1 {
    return index + 1
  }
  if nums[index+1] >= target {
    return index + 1
  }
  return index
}
