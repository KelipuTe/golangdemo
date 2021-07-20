package main

import "fmt"

func main() {
	iarrNums := []int{2, 7, 11, 15}
	iTarget := 9
	iarrRes := twoSum(iarrNums, iTarget)
	fmt.Println(iarrRes)
}

//两数之和
//从数组中找出两个整数，使得它们的和等于给定目标数，返回一个下标构成的数组
func twoSum(nums []int, target int) []int {
	//构造键为数和值为下标的map
	mapNum := make(map[int]int, len(nums))
	for iIndex1, iNum1 := range nums {
		mapNum[iNum1] = iIndex1
	}
	//用目标数依次和map的键表示的数做减法，检查减出来的数是不是map的键
	for iNum1, iIndex1 := range mapNum {
		iNum2 := target - iNum1
		_, bExist := mapNum[iNum2]
		if bExist {
			return []int{iIndex1, mapNum[iNum2]}
		}
	}

	return []int{-1, -1}
}
