package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

//给定整数数组nums和整数目标值target，找出和为目标值target的那两个整数，并返回它们的数组下标
//2<=nums.length<=10^4;-10^9<=nums[i]<=10^9;-10^9<=target<=10^9
//假设每种输入只会对应一个答案，但是，数组中同一个元素不能重复出现

//1-两数之和(1,15,167)
func twoSum(nums []int, target int) []int {
	//排序后双指针，或者构造map

	var mapNum map[int]int = make(map[int]int, len(nums))
	//构造，键为原数组值，值为原数组下标的map
	for iIndex1, iNum1 := range nums {
		mapNum[iNum1] = iIndex1
	}

	//遍历数组，得到遍历的值a1和a2=目标数减去a1，去map中查找有没有a2这个键
	for iNum1, iIndex1 := range mapNum {
		iNum2 := target - iNum1
		_, bExist := mapNum[iNum2]
		if bExist {
			return []int{iIndex1, mapNum[iNum2]}
		}
	}
	return []int{-1, -1}
}
