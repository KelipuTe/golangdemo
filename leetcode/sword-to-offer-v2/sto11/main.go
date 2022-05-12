package main

import "fmt"

// 剑指 Offer 11. 旋转数组的最小数字
// 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
// 给你一个可能存在 重复 元素值的数组 numbers ，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。
// 请返回旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为 1。

// 解题思路
// 旋转之后，数组有两种形态：1、不变。2、中间有一个断点，分隔左右两个升序序列。
// 取中点。如果中点等于右端点，那不知道拐点在哪，但是可以确定不是右端点，右端点左移一位。
// 如果中点小于右端点，说明拐点在左边，移动右端点；如果中点大于右端点，说明拐点在右边，移动左端点；

func main() {
  fmt.Println(minArray([]int{1, 5, 9}))
  fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
  fmt.Println(minArray([]int{3, 4, 5, 6, 7, 8, 9, 1, 2}))
  fmt.Println(minArray([]int{3, 1, 3}))
  fmt.Println(minArray([]int{3, 3, 1, 3}))
  fmt.Println(minArray([]int{1, 3, 3}))
}

func minArray(numbers []int) int {
  numLen := len(numbers)
  i, j := 0, numLen-1
  for i < j {
    mid := i + (j-i)>>1
    if numbers[mid] < numbers[j] {
      j = mid
    } else if numbers[mid] > numbers[j] {
      i = mid + 1
    } else {
      j--
    }
  }
  return numbers[i]
}
