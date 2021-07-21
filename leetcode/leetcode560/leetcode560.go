package main

import "fmt"

func main() {
	iarrNums := []int{-1, 1, 2, -2, 1, 1, 0}
	iTarget := 2
	// iarrNums := []int{1, 2, 3, 4}
	// iTarget := 9
	// iarrNums := []int{1, 1, 1}
	// iTarget := 2
	// iarrNums := []int{1, -1, 0}
	// iTarget := 0
	iRes := subarraySum(iarrNums, iTarget)
	fmt.Println(iRes)
}

//例：
//给定一个整数数组an和一个整数n，找到该数组中和为k的连续的子数组，返回子数组的个数。

//解：
//要找和为k的连续的子数组，也就是当i>j时，前i项的和减去前j项的和等于k
//需要构造一个键为前i项的和的值，值为前i项的和出现的次数的map
//基本逻辑就是，找前i项的和减去k的值，在前面出现过几次

//和为k的子数组
func subarraySum(nums []int, k int) int {
	iArrLen := len(nums)
	tiSum := 0
	mapSumNum := map[int]int{} //记录前i项的和为key出现过val次
	iRes := 0

	if iArrLen < 0 {
		return iRes
	}

	for ii := 0; ii < iArrLen; ii++ {
		tiSum += nums[ii]
		//直接累加出结果
		if tiSum == k {
			iRes++
		}
		//前i项的和减去前j项的和等于k
		_, bExist := mapSumNum[tiSum-k]
		if bExist {
			iRes += mapSumNum[tiSum-k]
		}
		//把前i项的和加到map里去，提供给后续循环使用
		mapSumNum[tiSum]++
	}

	return iRes
}
