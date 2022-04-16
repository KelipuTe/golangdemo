//全排列FullPermutation
package main

import "fmt"

func main() {
  var arr1Nums []int = []int{1, 2, 3, 4}
  var arr1NumsLen = len(arr1Nums)
  var arr1Visit []int = make([]int, arr1NumsLen)
  var arr1Res []int = []int{} //临时结果集
  quan2pai2lie4(arr1Nums, arr1Visit, arr1Res, arr1NumsLen)
}

func quan2pai2lie4(arr1Nums []int, arr1Visit []int, arr1Res []int, numsRemain int) {
  if numsRemain < 1 { //待处理的数字小于1的时候，说明全排列结束
    fmt.Println(arr1Res)
  }
  for index, isVisit := range arr1Visit { //遍历每一个数字
    if isVisit == 0 { //如果这个数字没有被访问过，就访问
      arr1Res = append(arr1Res, arr1Nums[index])                //结果集加一项
      arr1Visit[index] = 1                                      //标记为已访问
      quan2pai2lie4(arr1Nums, arr1Visit, arr1Res, numsRemain-1) //回溯
      arr1Res = arr1Res[:len(arr1Res)-1]                        //结果集去掉该元素
      arr1Visit[index] = 0                                      //重置该元素的访问状态
    }
  }
}
