package main

import (
	"fmt"
	"sort"
)

func main() {
	isli1Num := []int{1, 1, 2}
	permuteUnique(isli1Num)
	fmt.Println(isli2Res)
}

//给定一个可包含重复数字的序列nums ，按任意顺序返回所有不重复的全排列
//1<=nums.length<=8;-10<=nums[i]<=10

var isli2Res [][]int //结果集

//46-全排列II(39,46,47)
func permuteUnique(nums []int) [][]int {
	//回溯算法
	//沿用第46题的思路，区别在于本题这么做会产生大量重复的组合
	//解决这个问题，可以将数组排序，保证相同的数字都相邻
	//然后每次填入的元素是这个元素所在重复元素集合中从左往右第一个未被填过的
	//也就是对于一组重复的元素，填入顺序一定是像这样的，[0,0,0]=>[1,0,0]=>[1,1,0]=>[1,1,1]
	isli2Res = [][]int{}                            //初始化结果集
	var isli1Res []int = []int{}                    //每次回溯的结果
	var bsli1Visit []bool = make([]bool, len(nums)) //访问过的元素
	sort.Ints(nums)                                 //排序，目的是去除重复的组合
	hui2su4(nums, 0, isli1Res, bsli1Visit)
	return isli2Res
}

func hui2su4(nums []int, iIndex int, isli1Res []int, isli1Visit []bool) {
	iNumsLen := len(nums)

	if len(isli1Res) == iNumsLen {
		//回溯结果长度等于原数组长度，排列结束
		//这里要复制一份结果，因为切片是引用传递的，后续的回溯会影响已经存储的结果
		var tsli1Res []int = make([]int, iNumsLen)
		copy(tsli1Res, isli1Res)
		isli2Res = append(isli2Res, tsli1Res)
		return
	}

	for ii := 0; ii < iNumsLen; ii++ {
		if isli1Visit[ii] {
			continue
		}
		if ii > 0 && nums[ii] == nums[ii-1] && !isli1Visit[ii-1] {
			continue //左往右第一个未被填过的数字
		}
		isli1Visit[ii] = true                   //标记访问位置
		isli1Res = append(isli1Res, nums[ii])   //添加到回溯结果里
		hui2su4(nums, ii, isli1Res, isli1Visit) //进入下一层
		isli1Visit[ii] = false                  //重置访问位置
		isli1Res = isli1Res[:len(isli1Res)-1]   //从回溯结果里移除
	}
}
