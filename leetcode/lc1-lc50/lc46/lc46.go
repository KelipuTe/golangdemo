package main

import "fmt"

var isli2Res [][]int //结果集

func main() {
	isli1Num := []int{1, 2, 3, 4}
	permute(isli1Num)
	fmt.Println(isli2Res)
}

//46-全排列
func permute(nums []int) [][]int {
	//回溯算法

	isli2Res = [][]int{}
	var isli1Res []int = []int{}                  //回溯结果
	var isli1Visit []int = make([]int, len(nums)) //访问过的元素

	hui2su4(nums, isli1Res, isli1Visit)

	return isli2Res
}

func hui2su4(nums []int, isli1Res []int, isli1Visit []int) {
	iNumLen := len(nums)

	if iNumLen == len(isli1Res) {
		var tsli1Res []int = make([]int, iNumLen)
		copy(tsli1Res, isli1Res)
		isli2Res = append(isli2Res, tsli1Res)
		return
	}

	for ii := 0; ii < iNumLen; ii++ {
		if isli1Visit[ii] == 1 {
			continue
		}
		isli1Visit[ii] = 1
		isli1Res = append(isli1Res, nums[ii])
		hui2su4(nums, isli1Res, isli1Visit)
		isli1Visit[ii] = 0
		isli1Res = isli1Res[:len(isli1Res)-1]
	}
}
