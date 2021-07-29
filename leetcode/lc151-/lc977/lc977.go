package main

import "fmt"

func main() {
	isli1Res := sortedSquares([]int{-4, -1, 0, 3, 10})
	for iIndexZuo1 := 0; iIndexZuo1 < len(isli1Res); iIndexZuo1++ {
		fmt.Printf("%d,", isli1Res[iIndexZuo1])
	}
}

//有序数组的平方
func sortedSquares(nums []int) []int {
	//负数越小平方越大，正数越大平方越大
	//因为数组是有序的，所以最小的负数和最大的正数分别位于两侧
	//从两侧开始，将数字平方后比较，大的那个就是结果数组里大的那个

	iArrLen := len(nums)
	var isli1Nums []int = make([]int, iArrLen)
	var iIndexZuo1, iIndexYou4 int = 0, iArrLen - 1
	var iIndexRes int = iArrLen - 1 //存储结果的下标

	for iIndexZuo1 < iIndexYou4 {
		if nums[iIndexZuo1]*nums[iIndexZuo1] > nums[iIndexYou4]*nums[iIndexYou4] {
			isli1Nums[iIndexRes] = nums[iIndexZuo1] * nums[iIndexZuo1]
			iIndexRes--
			iIndexZuo1++
		} else {
			isli1Nums[iIndexRes] = nums[iIndexYou4] * nums[iIndexYou4]
			iIndexRes--
			iIndexYou4--
		}
	}

	return isli1Nums
}
