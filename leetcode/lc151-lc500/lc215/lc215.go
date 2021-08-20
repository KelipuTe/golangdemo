package main

import (
  "fmt"
)

func main() {
  fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
  fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

//给定整数数组nums和整数k，请返回数组中第k个最大的元素。
//请注意，需要找的是数组排序后的第k个最大的元素，而不是第k个不同的元素。
//1<=k<=nums.length<=10^4;-10^4<=nums[i]<=10^4

//快速排序
//排序后第k个最大元素，就是排序后倒数第k个元素，就是正数第n-k+1个元素
//用快速排序寻找排序后第k个元素的思路

//215-数组中的第k个最大元素
func findKthLargest(nums []int, k int) int {
  return kuai4su4pai2xu4(nums, 0, len(nums)-1, len(nums)-k)
}

//快速排序
func kuai4su4pai2xu4(nums []int, indexStart int, indexEnd int, indexK int) int {
  var indexStartTemp, indexEndTemp int = indexStart, indexEnd
  var midNum int = nums[indexStart]

  if indexStart >= indexEnd {
    return nums[indexStart]
  }

  for indexStartTemp < indexEndTemp {
    if nums[indexEndTemp] < midNum {
      nums[indexStartTemp], nums[indexEndTemp] = nums[indexEndTemp], nums[indexStartTemp]
      for indexStartTemp += 1; indexStartTemp < indexEndTemp; indexStartTemp++ {
        if nums[indexStartTemp] > midNum {
          nums[indexStartTemp], nums[indexEndTemp] = nums[indexEndTemp], nums[indexStartTemp]
          break
        }
      }
    }
    indexEndTemp--
  }

  //和快速排序不一样，这里只需要递归处理k存在的那部分
  if indexStartTemp < indexK {
    return kuai4su4pai2xu4(nums, indexStartTemp+1, indexEnd, indexK)
  } else if indexStartTemp > indexK {
    return kuai4su4pai2xu4(nums, indexStart, indexStartTemp, indexK)
  } else {
    return nums[indexStartTemp]
  }
}
