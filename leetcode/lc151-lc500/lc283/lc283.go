package main

import "fmt"

func main() {
  var sli1Nums []int = []int{0, 1, 0, 3, 12}
  moveZeroes(sli1Nums)
  fmt.Println(sli1Nums)
}

//给定一个数组nums，编写一个函数将所有0移动到数组的末尾，同时保持非零元素的相对顺序。
//必须在原数组上操作，不能拷贝额外的数组。尽量减少操作次数。

//数组，双指针
//把所有的非0元素移到前面，然后剩下的位置补0即可

//283-移动零
func moveZeroes(nums []int) {
  var numsLen int = len(nums)
  var indexRead, indexWrite int = 0, 0

  for indexRead < numsLen {
    if nums[indexRead] != 0 {
      nums[indexWrite] = nums[indexRead]
      indexWrite++
    }
    indexRead++
  }
  for indexWrite < numsLen {
    nums[indexWrite] = 0
    indexWrite++
  }
}
