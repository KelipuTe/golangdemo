package main

import "fmt"

func main() {
  fmt.Println(sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
  fmt.Println(sumOddLengthSubarrays([]int{1, 2}))
  fmt.Println(sumOddLengthSubarrays([]int{10, 11, 12}))
}

//给你一个正整数数组arr ，请你计算所有可能的奇数长度子数组的和。
//子数组定义为原数组中的一个连续子序列。请返回arr中所有奇数长度子数组的和 。

//数组，滑动窗口

//1588-所有奇数长度子数组的和
func sumOddLengthSubarrays(arr []int) int {
  var arrLen int = len(arr)
  var sumRes int = 0 //结果
  var sumLen int = 1 //窗口长度

  for sumLen <= arrLen {
    sumResTemp := 0 //每个窗口内元素的和
    index := 0
    //初始化窗口
    for index < sumLen {
      sumResTemp += arr[index]
      index++
    }
    sumRes += sumResTemp
    //滑动窗口
    for index < arrLen {
      sumResTemp += arr[index]
      sumResTemp -= arr[index-sumLen]
      sumRes += sumResTemp
      index++
    }
    sumLen += 2 //只算奇数，所以窗口长度+2
  }

  return sumRes
}
