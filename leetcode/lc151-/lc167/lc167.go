package main

import "fmt"

func main() {
	iarrRes := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(iarrRes)
}

//两数之和
//从有序数组an中找出两个整数[a1,a2]，使得它们的和等于给定目标数n，数组元素会重复
func twoSum(numbers []int, target int) []int {
	//这里由于数组元素会重复，所以不能用两数之和（leetcode1）的思路
	//这里可以参考三数之和（leetcode15）的逻辑，区别在于只需要在基数a1右侧找一个数a2=n-a1就行了
	//由于数组是有序的，所以可以用二分查找找这个数

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
