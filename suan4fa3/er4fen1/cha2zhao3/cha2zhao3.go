//二分查找，BinarySearch
//二分查找的条件比较苛刻，需要数组格式和排好序的数据
package main

import (
  "fmt"
)

func main() {
  fmt.Println(er4fen1e([]int{2, 4, 6, 8, 10, 12, 14, 16}, 8))

  fmt.Println(er4fen1gt([]int{2, 4, 6, 8, 10, 12, 14, 16}, 10))
  fmt.Println(er4fen1gt([]int{2, 2, 4, 4, 4, 4, 4, 8, 8}, 4))

  fmt.Println(er4fen1lt([]int{2, 4, 6, 8, 10, 12, 14, 16}, 10))
  fmt.Println(er4fen1lt([]int{2, 2, 4, 4, 4, 4, 4, 8, 8}, 4))
}

//二分查找查找升序数组中，等于目标数的元素的下标
func er4fen1e(nums []int, target int) int {
  var numsLen int = len(nums)
  var indexLeft, indexRight int = 0, numsLen - 1

  for indexLeft <= indexRight { //注意这里是小于等于
    var indexMid int = indexLeft + (indexRight-indexLeft)>>1 //取中，这么写可以防止indexLeft+indexRight溢出

    if nums[indexMid] > target { //中间数大于目标数，目标数在左半边
      indexRight = indexMid - 1 //收缩右边界，注意-1
    } else if nums[indexMid] < target { //中间数小于目标数，目标数在右半边
      indexLeft = indexMid + 1 //收缩左边界，注意+1
    } else { //命中目标数
      return indexMid
    }
  }
  return -1
}

//二分查找，查找升序数组中，大于目标数的最小的元素的下标
func er4fen1gt(nums []int, target int) int {
  var numsLen int = len(nums)
  var indexLeft, indexRight int = 0, numsLen - 1
  var indexStart = numsLen //结果位置，初始化为数组长度，如果没有变过，说明目标数比数组中任意一个元素都大

  for indexLeft <= indexRight {
    var indexMid int = indexLeft + (indexRight-indexLeft)>>1

    if nums[indexMid] > target {
      indexRight = indexMid - 1
      indexStart = indexMid //需要收缩右边界，则说明大于目标数的最小的元素在中间数位置或者中间数左侧位置
    } else {
      indexLeft = indexMid + 1
    }
  }
  return indexStart
}

//二分查找，查找升序数组中，小于目标数的最大的元素的下标
func er4fen1lt(nums []int, target int) int {
  var numsLen int = len(nums)
  var indexLeft, indexRight int = 0, numsLen - 1
  var indexStart = -1 //结果位置，初始化为-1，如果没有变过，说明目标数比数组中任意一个元素都小

  for indexLeft <= indexRight {
    indexMid := indexLeft + (indexRight-indexLeft)>>1

    if nums[indexMid] < target {
      indexLeft = indexMid + 1
      indexStart = indexMid //需要收缩左边界，则说明小于目标数的最大的元素在中间数位置或者中间数右侧位置
    } else {
      indexRight = indexMid - 1
    }
  }
  return indexStart
}
