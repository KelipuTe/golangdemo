package main

import "fmt"

func main() {
  fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

//给定一个已按照升序排列的整数数组numbers，请从数组中找出两个数满足相加之和等于目标数target。
//函数应该以长度为2的整数数组的形式返回这两个数的下标值。
//numbers的下标从1开始计数，所以答案数组应当满足1<=answer[0]<answer[1]<=numbers.length。
//假设每个输入只对应唯一的答案，而且不可以重复使用相同的元素。

//数组，双指针，二分查找
//基本思路和两数之和一样，遍历数组，得到遍历的值a1，计算另一个数a2=目标数-a1。
//然后，查找有没有数组中有没有a2，由于数组是升序的，所以可以用二分查找找这个数。
//2<=numbers.length<=3*104;-1000<=numbers[i]<=1000;
//numbers按递增顺序排列;-1000<=target<=1000;仅存在一个有效答案;
//这题依然可以沿用两数之和使用的哈希表的思路，但是哈希表的思路没有使用到数组有序的性质。

//167-两数之和II-输入有序数组
func twoSum(numbers []int, target int) []int {
  var numbersLen int = len(numbers)

  for index := 0; index < numbersLen; index++ {
    indexLeft, indexRight := index+1, numbersLen-1
    for indexLeft <= indexRight {
      indexMid := indexLeft + (indexRight-indexLeft)>>1
      if numbers[indexMid] > target-numbers[index] {
        indexRight = indexMid - 1
      } else if numbers[indexMid] < target-numbers[index] {
        indexLeft = indexMid + 1
      } else {
        return []int{index + 1, indexMid + 1}
      }
    }
  }

  return []int{-1, -1}
}
