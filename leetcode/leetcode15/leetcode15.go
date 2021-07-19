package main

import (
	"fmt"
	"sort"
)

func main() {
	// iarrNums := []int{0, 0, 0}
	// iarrNums := []int{-1, 0, 1, 2, -1, -4}
	iarrNums := []int{-2, 0, 3, -1, 4, 0, 3, 4, 1, 1, 1, -3, -5, 4, 0}
	iarrRes := threeSum(iarrNums)
	fmt.Println(iarrRes)
}

//三数之和
//从数组中找出三个整数，使得它们的和等于给定目标数
//返回一个数组，数组的每个元素是一组合符合条件的数构成的数组，每个组合只记一次
func threeSum(nums []int) [][]int {
	iTarget := 0
	iarr2sliRes := make([][]int, 0)
	iArrLen := len(nums)
	if iArrLen < 3 {
		return [][]int{}
	}
	sort.Ints(nums) //排序

	//固定一个基数，然后使用双指针遍历剩余的数组
	for iIndex1 := 0; iIndex1 < iArrLen-2; iIndex1++ {
		//跳过一样的数
		if iIndex1 > 0 && nums[iIndex1] == nums[iIndex1-1] {
			continue
		}
		iIndex2 := iIndex1 + 1
		iIndex3 := iArrLen - 1
		for iIndex2 < iIndex3 {
			if nums[iIndex1]+nums[iIndex2]+nums[iIndex3] > iTarget {
				//如果三数之和小于目标数，就需要左指针右移，提供一个更大的数
				for iIndex2 < iIndex3 && nums[iIndex3-1] == nums[iIndex3] {
					iIndex3--
				}
				iIndex3--
			} else if nums[iIndex1]+nums[iIndex2]+nums[iIndex3] < iTarget {
				//如果三数之和大于目标数，就需要右指针左移，提供一个更小的数
				for iIndex2 < iIndex3 && nums[iIndex2] == nums[iIndex2+1] {
					iIndex2++
				}
				iIndex2++
			} else {
				//找到符合条件的一组数，因为不能重复，所以两边的指针都可以移动一位
				iarr2sliRes = append(iarr2sliRes, []int{nums[iIndex1], nums[iIndex2], nums[iIndex3]})
				for iIndex2 < iIndex3 && nums[iIndex2] == nums[iIndex2+1] {
					iIndex2++
				}
				iIndex2++
				for iIndex2 < iIndex3 && nums[iIndex3-1] == nums[iIndex3] {
					iIndex3--
				}
				iIndex3--
			}
		}
	}

	return iarr2sliRes
}
