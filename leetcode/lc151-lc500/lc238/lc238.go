package main

import (
  "fmt"
)

func main() {
  fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
  fmt.Println(productExceptSelfPuls([]int{1, 2, 3, 4}))
}

//给你一个长度为n的整数数组nums，其中n>1，返回输出数组output，
//其中output[i]等于nums 中除nums[i]之外其余各元素的乘积。
//题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在32位整数范围内。
//请不要使用除法，且在O(n)时间复杂度内完成此题。
//进阶：在常数空间复杂度内完成这个题目吗？（出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

//使用两个数组分别存储数组中每个元素，左侧所有元素的乘积和右侧所有元素的乘积
//然后就可以通过下标计算数组中每个元素，左侧所有元素的乘积和右侧所有元素的乘积的乘积

//进阶解法可以使用结果集暂存左侧所有元素的乘积，然后用一个整型变量存储右侧所有元素的乘积
//从右开始往左算，每算一个元素，需要额外维护右侧所有元素的乘积，然后继续

//238-除自身以外数组的乘积
func productExceptSelfPuls(nums []int) []int {
  var numsLen = len(nums)
  var you4cheng2ji1 int //右侧所有元素的乘积的乘积
  var sli1Res []int

  sli1Res = make([]int, numsLen)

  //用结果集暂存左侧所有元素的乘积
  sli1Res[0] = 0
  sli1Res[1] = nums[0]
  for index := 2; index < numsLen; index++ {
    sli1Res[index] = sli1Res[index-1] * nums[index-1]
  }

  you4cheng2ji1 = nums[numsLen-1]
  for index := numsLen - 2; index > 0; index-- {
    sli1Res[index] = sli1Res[index] * you4cheng2ji1
    you4cheng2ji1 *= nums[index]
  }
  sli1Res[0] = you4cheng2ji1

  return sli1Res
}

func productExceptSelf(nums []int) []int {
  var numsLen = len(nums)
  var sli1Zuo3, sli1You4 []int //每个数组元素，左侧所有元素的乘积和右侧所有元素的乘积
  var sli1Res []int

  sli1Zuo3 = make([]int, numsLen)
  sli1You4 = make([]int, numsLen)
  sli1Res = make([]int, numsLen)

  sli1Zuo3[0] = nums[0]
  for index := 1; index < numsLen; index++ {
    sli1Zuo3[index] = sli1Zuo3[index-1] * nums[index]
  }
  sli1You4[numsLen-1] = nums[numsLen-1]
  for index := numsLen - 2; index > -1; index-- {
    sli1You4[index] = sli1You4[index+1] * nums[index]
  }

  sli1Res[0] = sli1You4[1]
  sli1Res[numsLen-1] = sli1Zuo3[numsLen-2]
  for index := 1; index < numsLen-1; index++ {
    sli1Res[index] = sli1Zuo3[index-1] * sli1You4[index+1]
  }

  return sli1Res
}
