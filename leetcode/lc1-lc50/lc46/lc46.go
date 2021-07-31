package main

import "fmt"

func main() {
	isli1Num := []int{1, 2, 3, 4}
	permute(isli1Num)
	fmt.Println(isli2Res)
}

//给定一个不含重复数字的数组nums，返回其所有可能的全排列
//1<=nums.length<=6;-10<=nums[i]<=10

var isli2Res [][]int //结果集

//46-全排列(39,46,47)
func permute(nums []int) [][]int {
	//回溯算法
	//排列时元素不能重复，所以需要标记访问过的元素
	isli2Res = [][]int{}                            //初始化结果集
	var isli1Res []int = []int{}                    //每次回溯的结果
	var isli1Visit []bool = make([]bool, len(nums)) //访问过的元素
	hui2su4(nums, isli1Res, isli1Visit)
	return isli2Res
}

func hui2su4(nums []int, isli1Res []int, isli1Visit []bool) {
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
		isli1Visit[ii] = true                 //标记访问位置
		isli1Res = append(isli1Res, nums[ii]) //添加到回溯结果里
		hui2su4(nums, isli1Res, isli1Visit)   //进入下一层
		isli1Visit[ii] = false                //重置访问位置
		isli1Res = isli1Res[:len(isli1Res)-1] //从回溯结果里移除
	}
}
