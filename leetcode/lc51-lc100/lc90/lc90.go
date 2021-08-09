package main

import (
  "fmt"
  "sort"
)

func main() {
  fmt.Println(subsetsWithDup([]int{1, 2, 2, 2, 3}))
}

//给你一个整数数组nums，其中可能包含重复元素，返回该数组所有可能的子集（幂集）
//解集不能包含重复的子集。返回的解集中，子集可以按任意顺序排列
// 1<=nums.length<=10;-10<=nums[i]<=10

var isli2Res [][]int //结果集

//回溯
//沿用第78题的思路，每次回溯都有两种选择，要自己或者跳过自己，区别在于本题这么做会产生大量重复的组合
//这里再借助第47题的思路，可以将数组排序，保证相同的数字都相邻
//然后每次决定要自己时，自己必须是所在重复元素集合中从左往右第一个未被填过的
//也就是对于一组重复的元素，填入顺序一定是像这样的，[0,0,0]=>[1,0,0]=>[1,1,0]=>[1,1,1]

//90-子集II(39,46,47,78,90)
func subsetsWithDup(nums []int) [][]int {
  isli2Res = [][]int{}                            //初始化结果集
  var isli1Res []int = []int{}                    //每次回溯的结果
  var bsli1Visit []bool = make([]bool, len(nums)) //访问过的元素
  sort.Ints(nums)                                 //排序，目的是去除重复的组合
  hui2su4(nums, 0, len(nums), isli1Res, bsli1Visit)
  return isli2Res
}

// iIndex 本次回溯的元素下标
func hui2su4(nums []int, iIndex int, iNumsLen int, isli1Res []int, isli1Visit []bool) {
  if iIndex == iNumsLen {
    //所有的元素都判断一次后结束
    var tsli1Res []int = make([]int, len(isli1Res))
    copy(tsli1Res, isli1Res)
    isli2Res = append(isli2Res, tsli1Res)
    return
  }
  //要自己
  if iIndex > 0 && nums[iIndex] == nums[iIndex-1] && !isli1Visit[iIndex-1] {
    //不是重复元素集合中左往右第一个未被填过的数字，这时要自己不满足条件
  } else {
    isli1Visit[iIndex] = true
    isli1Res = append(isli1Res, nums[iIndex])
    hui2su4(nums, iIndex+1, iNumsLen, isli1Res, isli1Visit)
    isli1Visit[iIndex] = false
    isli1Res = isli1Res[:len(isli1Res)-1]
  }
  //跳过自己，直接下一个
  hui2su4(nums, iIndex+1, iNumsLen, isli1Res, isli1Visit)
}
