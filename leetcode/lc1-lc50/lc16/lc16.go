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

//例：
//从数组中找出三个整数，让它们的和最接近给定的目标数，返回最接近的那个和

//解：
//基本思路同三数之和（leetcode15）
//区别在于每次遍历需要记录当前三数之和和目标数的差值并和目前最小的差值比较

//最接近的三数之和
func threeSumClosest(nums []int, target int) int {
	iArrLen := len(nums)
	iClosestSum := 65535                //最接近的三数之和，初始化一个临界值
	iClosestLen := iClosestSum - target //最接近的三数之和和目标数的差值

	if iArrLen < 3 {
		return iClosestSum
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
