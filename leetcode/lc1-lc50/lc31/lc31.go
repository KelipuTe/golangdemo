package main

import (
	"fmt"
	"sort"
)

func main() {
	isli1Num := []int{1, 2, 3, 4, 3, 2, 1}
	// isli1Num := []int{3, 2, 1}
	// isli1Num := []int{1, 2}
	nextPermutation(isli1Num)
	fmt.Println(isli1Num)
}

//算法需要将给定数字序列重新排列成字典序中下一个更大的排列
//如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）
//必须原地修改，只允许使用额外常数空间

//31-下一个排列
func nextPermutation(nums []int) {
	//从后往前找，找到第一个数n[a]，让n[a]比n[a+1]小，如果没找到，说明数组是降序的，已经是最大的数字了
	//再从后往前找，找到第一个比n[a]大的数n[b]，交换n[a]和n[b]，再把交换后n[b]后面的序列反序
	//比如：[4,5,3,6,4,2,1]，n[a]=n[2]=3，n[b]=n[4]=4，交换后：[4,5,4,6,3,2,1]，反序后：[4,5,4,1,2,3,6]

	iNumsLen := len(nums)

	//后往前找n[a]
	iIndex1 := -1
	for ii := iNumsLen - 2; ii >= 0; ii-- {
		if nums[ii] < nums[ii+1] {
			iIndex1 = ii
			break
		}
	}
	if iIndex1 == -1 {
		sort.Ints(nums)
		return
	}
	//后往前找n[b]
	iIndex2 := -1
	for ii := iNumsLen - 1; ii > iIndex1; ii-- {
		if nums[ii] > nums[iIndex1] {
			iIndex2 = ii
			break
		}
	}
	nums[iIndex1], nums[iIndex2] = nums[iIndex2], nums[iIndex1]
	//反序
	if iIndex1 < iNumsLen-2 {
		ii := iIndex1 + 1
		ij := iNumsLen - 1
		for ii < ij {
			nums[ii], nums[ij] = nums[ij], nums[ii]
			ii++
			ij--
		}
	}
}
