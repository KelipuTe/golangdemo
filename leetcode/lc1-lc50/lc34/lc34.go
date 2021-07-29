package main

import "fmt"

func main() {
	// fmt.Println(searchRange([]int{4, 4, 5, 5, 7, 7, 8, 8, 8, 9, 9, 9, 10}, 5))
	fmt.Println(searchRange([]int{1}, 0))
}

//34、在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	//二分查找
	iNumsLen := len(nums)
	if iNumsLen < 1 {
		return []int{-1, -1}
	}

	//找到最小index的target
	iLeft, iRight := 0, iNumsLen-1
	for iLeft < iRight {
		//右边往左收
		iMid := iLeft + (iRight-iLeft-1)/2
		if nums[iMid] >= target {
			iRight = iMid
		} else {
			iLeft = iMid + 1
		}
	}
	iFirst := iRight

	//找到最大index的target
	iLeft, iRight = 0, iNumsLen-1
	for iLeft < iRight {
		//左边往有收
		iMid := iLeft + (iRight-iLeft+1)/2
		if nums[iMid] <= target {
			iLeft = iMid
		} else {
			iRight = iMid - 1
		}
	}
	iLast := iLeft

	if iFirst > iLast {
		return []int{-1, -1}
	}
	return []int{iFirst, iLast}
}
