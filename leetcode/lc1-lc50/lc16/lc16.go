package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{1, 1, -1, -1, 3}, 9))
}

//给定整数的数组nums和一个目标值target
//找出nums中的三个整数，使得它们的和与target最接近，返回这三个数的和
//3<=nums.length<=10^3;-10^3<=nums[i]<=10^3;-10^4<=target<=10^4
//假定每组输入只存在唯一答案

//16-最接近的三数之和(15,16)
func threeSumClosest(nums []int, target int) int {
	//基本思路和第15题三数之和相同
	//区别在于每次遍历需要记录当前三数之和和目标数的差值并和目前最小的差值比较

	iNumsLen := len(nums)
	iClosestSum := 65535                //最接近的三数之和，初始化一个极大临界值
	iClosestLen := iClosestSum - target //最接近的三数之和和目标数的差值

	if iNumsLen < 3 {
		return 65535
	}

	sort.Ints(nums) //排序
	for iIndex1 := 0; iIndex1 < iNumsLen-2; iIndex1++ {
		if iIndex1 > 0 && nums[iIndex1] == nums[iIndex1-1] {
			continue
		}
		iIndex2, iIndex3 := iIndex1+1, iNumsLen-1
		for iIndex2 < iIndex3 {
			tiSum := nums[iIndex1] + nums[iIndex2] + nums[iIndex3]
			if tiSum > target {
				//三数之和大于目标数
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
				//三数之和小于目标数
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
				//三数之和等于目标数
				iClosestSum = target
				iClosestLen = 0
				break
			}
		}
	}

	return iClosestSum
}
