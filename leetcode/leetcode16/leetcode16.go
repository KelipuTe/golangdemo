package main

import (
	"fmt"
	"sort"
)

func main() {
	// iarrNums := []int{-1, 2, 1, -4}
	iarrNums := []int{1, 1, -1, -1, 3}
	iTarget := 1
	iarrRes := threeSumClosest(iarrNums, iTarget)
	fmt.Println(iarrRes)
}

//最接近的三数之和
//从数组中找出三个整数，让它们的和最接近给定的目标数，返回最接近的那个和
func threeSumClosest(nums []int, target int) int {
	iClosestSum := 65535                //最接近的三数之和，初始化一个临界值
	iClosestLen := iClosestSum - target //最接近的三数之和和目标数的差值
	iArrLen := len(nums)

	if iArrLen < 3 {
		return iClosestSum
	}
	//排序
	sort.Ints(nums)
	//固定一个基数，然后使用双指针遍历剩余的数组
	for iIndex1 := 0; iIndex1 < iArrLen-2; iIndex1++ {
		//跳过一样的数
		if iIndex1 > 0 && nums[iIndex1] == nums[iIndex1-1] {
			continue
		}
		iIndex2 := iIndex1 + 1
		iIndex3 := iArrLen - 1
		for iIndex2 < iIndex3 {
			tiSum := nums[iIndex1] + nums[iIndex2] + nums[iIndex3]
			if tiSum > target {
				//三数之和小于目标数
				if tiSum-target < iClosestLen {
					iClosestSum = tiSum
					if tiSum-target > 0 {
						iClosestLen = tiSum - target
					} else {
						iClosestLen = target - tiSum
					}
				}
				//左指针右移，提供一个更大的数
				for iIndex2 < iIndex3 && nums[iIndex3-1] == nums[iIndex3] {
					iIndex3--
				}
				iIndex3--
			} else if tiSum < target {
				//三数之和大于目标数
				if target-tiSum < iClosestLen {
					iClosestSum = tiSum
					if tiSum-target > 0 {
						iClosestLen = tiSum - target
					} else {
						iClosestLen = target - tiSum
					}
				}
				//右指针左移，提供一个更小的数
				for iIndex2 < iIndex3 && nums[iIndex2] == nums[iIndex2+1] {
					iIndex2++
				}
				iIndex2++
			} else {
				//完全命中
				iClosestSum = target
				iClosestLen = 0
				break
			}
		}
	}

	return iClosestSum
}
