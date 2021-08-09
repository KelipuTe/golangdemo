package main

import "fmt"

func main() {
  // fmt.Println(subsets([]int{1, 2, 3}))
  fmt.Println(subsets([]int{1}))
}

//给一个整数数组nums ，数组中的元素互不相同。返回该数组所有可能的子集（幂集）
//解集不能包含重复的子集。可以按任意顺序返回解集
//1<=nums.length<=10;-10<=nums[i]<=10;nums中的所有元素互不相同

//回溯
//每次回溯都有两种选择，要自己或者跳过自己

var isli2Res [][]int //结果集

//78-子集(39,46,47,78,90)
func subsets(nums []int) [][]int {
  isli2Res = [][]int{}         //初始化结果集
  var isli1Res []int = []int{} //每次回溯的结果
  hui2su4(nums, 0, len(nums), isli1Res)
  return isli2Res
}

// iIndex 本次回溯的元素下标
func hui2su4(nums []int, iIndex int, iNumsLen int, isli1Res []int) {
  if iIndex >= iNumsLen {
    //所有的元素都判断一次后结束
    var tsli1Res []int = make([]int, len(isli1Res))
    copy(tsli1Res, isli1Res)
    isli2Res = append(isli2Res, tsli1Res)
    return
  }

  //要自己
  isli1Res = append(isli1Res, nums[iIndex])
  hui2su4(nums, iIndex+1, iNumsLen, isli1Res)
  isli1Res = isli1Res[:len(isli1Res)-1]
  //跳过自己，直接下一个
  hui2su4(nums, iIndex+1, iNumsLen, isli1Res)
}
