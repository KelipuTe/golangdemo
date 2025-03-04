package main

import "fmt"

func main() {
  fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
  fmt.Println(twoSum([]int{3, 2, 4}, 6))
  fmt.Println(twoSum([]int{3, 3}, 6))
}

//给定一个整数数组nums和一个整数目标值target，在该数组中找出和为目标值target的那两个整数，并返回它们的数组下标。
//假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。可以按任意顺序返回答案。
//2<=nums.length<=10^4;-10^9<=nums[i]<=10^9;-10^9<=target<=10^9;只会存在一个有效答案;
//进阶：实现一个时间复杂度小于O(n^2)的算法。

//数组，哈希表（或，排序，双指针）
//构造一个map，键为原数组值，值为原数组下标。
//遍历数组，得到遍历的值a1，计算另一个数a2=目标数-a1。
//然后，去map中查找有没有键为a2的，注意元素不能重复使用。
//如果先完整的构造map，会存在无法处理3+3=6这样，两个数字是一样的问题场景。
//所以，可以一边遍历一边构造map，这样前一个3会被写入map，
//这时遍历到后一个3时，直接就可以匹配出结果，不会有上面的问题。

//1-两数之和
func twoSum(nums []int, target int) []int {
  var mapNums map[int]int = map[int]int{}
  for index, num := range nums {
    if index2, canFind := mapNums[target-num]; canFind {
      return []int{index2, index}
    }
    mapNums[num] = index
  }

  return []int{-1, -1}
}
