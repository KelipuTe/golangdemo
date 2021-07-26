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

//下一个排列
//给定一个数组，数组中的元素组成一个数字，
//将数组重新排列，找到可以排列出的，比这个数字大的第一个数字，如果没有就将数组升序排序
func nextPermutation(nums []int) {
	//从后往前找，找到第一个数n[a]，让n[a]比n[a+1]小，如果没找到，说明数组是降序的，已经是最大的数字了
	//再从后往前找，找到第一个比n[a]大的数n[b]，交换n[a]和n[b]，再把交换后n[b]后面的序列反序
	//比如：[4,5,3,6,4,2,1]，n[a]=n[2]=3，n[b]=n[4]=4，交换后：[4,5,4,6,3,2,1]，反序后：[4,5,4,1,2,3,6]

	iNumsLen := len(nums)

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

	iIndex2 := -1
	for ii := iNumsLen - 1; ii > iIndex1; ii-- {
		if nums[ii] > nums[iIndex1] {
			iIndex2 = ii
			break
		}
	}

	tiNum := nums[iIndex1]
	nums[iIndex1] = nums[iIndex2]
	nums[iIndex2] = tiNum

	if iIndex1 < iNumsLen-2 {
		ii := iIndex1 + 1
		ij := iNumsLen - 1
		for ii < ij {
			tiNum := nums[ii]
			nums[ii] = nums[ij]
			nums[ij] = tiNum
			ii++
			ij--
		}
	}
}
