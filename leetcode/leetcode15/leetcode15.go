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

//例：
//从数组an中找出三个整数[a1,a2,a3]，使得它们的和等于给定目标数n，数组元素会重复

//解：
//先对原数组进行升序排序
//从左往右，依次固定基数a1，在右侧的剩下的数组元素中找到两个整数a2和a3使得它们的和等于n-a1
//使用双指针，a2从i1+1开始往右，a3从in-1开始往左
//如果a2+a3>n-a1，则需要提供一个更小的和，所以a3左移，反之a2右移

//三数之和
func threeSum(nums []int) [][]int {
	iArrLen := len(nums)
	iTarget := 0
	isli2Res := make([][]int, 0)

	if iArrLen < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	for iIndex1 := 0; iIndex1 < iArrLen-2; iIndex1++ {
		//跳过一样的数
		if iIndex1 > 0 && nums[iIndex1] == nums[iIndex1-1] {
			continue
		}
		iIndex2 := iIndex1 + 1
		iIndex3 := iArrLen - 1
		for iIndex2 < iIndex3 {
			if nums[iIndex1]+nums[iIndex2]+nums[iIndex3] > iTarget {
				//三数之和小于目标数，左指针右移
				for iIndex2 < iIndex3 && nums[iIndex3-1] == nums[iIndex3] {
					iIndex3--
				}
				iIndex3--
			} else if nums[iIndex1]+nums[iIndex2]+nums[iIndex3] < iTarget {
				//三数之和大于目标数，右指针左移
				for iIndex2 < iIndex3 && nums[iIndex2] == nums[iIndex2+1] {
					iIndex2++
				}
				iIndex2++
			} else {
				//找到符合条件的一组数，因为不能重复，所以两边的指针都可以移动一位
				isli2Res = append(isli2Res, []int{nums[iIndex1], nums[iIndex2], nums[iIndex3]})
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

	return isli2Res
}
