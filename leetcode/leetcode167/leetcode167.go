package main

import "fmt"

func main() {
	iarrNums := []int{2, 7, 11, 15}
	iTarget := 9
	iarrRes := twoSum(iarrNums, iTarget)
	fmt.Println(iarrRes)
}

//两数之和
//从有序数组中找出两个整数，使得它们的和等于给定目标数，返回一个下标构成的数组
func twoSum(numbers []int, target int) []int {
	//用目标数依次和数组元素做减法，由于数组是有序的，所以可以用二分查找，查找减出来的数在哪里
	iArrLen := len(numbers)
	for ii := 0; ii < iArrLen; ii++ {
		iLow := ii + 1
		iHigh := iArrLen - 1
		for iLow <= iHigh {
			iMid := (iHigh-iLow)/2 + iLow
			if numbers[iMid] == target-numbers[ii] {
				return []int{ii + 1, iMid + 1}
			} else if numbers[iMid] > target-numbers[ii] {
				iHigh = iMid - 1
			} else {
				iLow = iMid + 1
			}
		}
	}
	return []int{-1, -1}
}
