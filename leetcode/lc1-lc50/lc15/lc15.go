package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(threeSum([]int{0, 0, 0}))
	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-2, 0, 3, -1, 4, 0, 3, 4, 1, 1, 1, -3, -5, 4, 0}))
}

//给定整数数组nums，判断nums中是否存在三个元素a，b，c，使得a+b+c=0
//找出所有和为0且不重复的三元组，答案中不可以包含重复的三元组
//0<=nums.length<=3000;-10^5<=nums[i]<=10^5

//15-三数之和(1,15,167)(15,16)
func threeSum(nums []int) [][]int {
	//排序后双指针
	//从左往右，依次固定基数a1，在右侧的剩下的数组元素中找到两个整数a2和a3，使得它们的和等于0-a1
	//使用双指针，下标i2(a2)从a1右边开始往右，下标i3(a3)从数组尾部开始往左
	//如果a2+a3>n-a1，则需要提供一个更小的和，所以a3左移，反之a2右移
	//因为答案中不可以包含重复的三元组，所以遇到连续的数字可以跳过

	iNumsLen := len(nums)
	if iNumsLen < 3 {
		return [][]int{}
	}

	var isli2Res [][]int = make([][]int, 0) //结果

	sort.Ints(nums) //排序
	for iIndex1 := 0; iIndex1 < iNumsLen-2; iIndex1++ {
		if iIndex1 > 0 && nums[iIndex1] == nums[iIndex1-1] {
			continue
		}
		iIndex2, iIndex3 := iIndex1+1, iNumsLen-1
		for iIndex2 < iIndex3 {
			if nums[iIndex1]+nums[iIndex2]+nums[iIndex3] > 0 {
				//a2+a3>n-a1，a3左移
				for iIndex2 < iIndex3 && nums[iIndex3-1] == nums[iIndex3] {
					iIndex3--
				}
				iIndex3--
			} else if nums[iIndex1]+nums[iIndex2]+nums[iIndex3] < 0 {
				//a2+a3<n-a1，a2右移
				for iIndex2 < iIndex3 && nums[iIndex2] == nums[iIndex2+1] {
					iIndex2++
				}
				iIndex2++
			} else {
				//a2+a3=n-a1，找到符合条件的一组数
				isli2Res = append(isli2Res, []int{nums[iIndex1], nums[iIndex2], nums[iIndex3]})
				//因为不能重复，所以两边的指针都可以移动一位
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
