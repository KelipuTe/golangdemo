package main

import "fmt"

func main() {
	// fmt.Println(searchRange([]int{4, 4, 5, 5, 7, 7, 8, 8, 8, 9, 9, 9, 10}, 4))
	// fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	// fmt.Println(searchRange([]int{2, 2}, 1))
	fmt.Println(searchRange([]int{1}, 1))
}

//给定升序排列的整数数组和一个目标值target，找出给定目标值在数组中的开始位置和结束位置
//0<=nums.length<=10^5，-10^9<=nums[i]<=10^9，-10^9<=target<=10^9

//34-在排序数组中查找元素的第一个和最后一个位置(34,278)
func searchRange(nums []int, target int) []int {
	//二分查找
	iNumsLen := len(nums)

	if iNumsLen < 1 {
		return []int{-1, -1}
	}
	if iNumsLen == 1 && nums[0] != target {
		return []int{-1, -1}
	}

	//二分查找，查找大于目标数的最小的元素的下标
	iIndex1 := iNumsLen
	il1, ir1 := 0, iNumsLen-1
	for il1 <= ir1 {
		im1 := il1 + (ir1-il1)>>1
		if nums[im1] > target {
			ir1 = im1 - 1
			iIndex1 = im1
		} else {
			il1 = im1 + 1
		}
	}

	//二分查找，查找小于目标数的最大的元素的下标
	iIndex2 := -1
	il2, ir2 := 0, iNumsLen-1
	for il2 <= ir2 {
		im2 := il2 + (ir2-il2)>>1
		if nums[im2] < target {
			il2 = im2 + 1
			iIndex2 = im2
		} else {
			ir2 = im2 - 1
		}
	}

	if iIndex1-1 >= iIndex2+1 && nums[iIndex1-1] == nums[iIndex2+1] {
		return []int{iIndex2 + 1, iIndex1 - 1}
	}

	return []int{-1, -1}
}
