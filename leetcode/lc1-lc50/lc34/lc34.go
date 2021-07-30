package main

import "fmt"

func main() {
	// fmt.Println(searchRange([]int{4, 4, 5, 5, 7, 7, 8, 8, 8, 9, 9, 9, 10}, 7))
	// fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{2, 2}, 1))
	// fmt.Println(searchRange([]int{1}, 1))
}

//34、在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	//二分查找
	iNumsLen := len(nums)
	if iNumsLen < 1 {
		return []int{-1, -1}
	}
	if iNumsLen == 1 && nums[0] != target {
		return []int{-1, -1}
	}

	//不等的时候，和二分查找一样处理
	//等于的时候，右边界往左边压，找到最小index的target
	iLeft, iRight := 0, iNumsLen-1
	for iLeft < iRight {
		iMid := iLeft + (iRight-iLeft)>>1
		if nums[iMid] > target {
			iRight = iMid - 1
		} else if nums[iMid] == target {
			//要让iMid不会卡死
			if iMid > 0 {
				if nums[iMid-1] == target {
					iRight = iMid - 1
				} else {
					iRight = iMid
					iLeft++
				}
			} else if iMid == 0 {
				iRight = iMid
			}
		} else {
			iLeft = iMid + 1
		}
	}
	iFirst := iRight

	//不等的时候，和二分查找一样处理
	//等于的时候，左边界往右边压，找到最大index的target
	iLeft, iRight = 0, iNumsLen-1
	for iLeft < iRight {
		iMid := iLeft + (iRight-iLeft)>>1
		if nums[iMid] < target {
			iLeft = iMid + 1
		} else if nums[iMid] == target {
			//要让iMid不会卡死
			if iMid < iNumsLen-1 {
				if nums[iMid+1] == target {
					iLeft = iMid + 1
				} else {
					iLeft = iMid
					iRight--
				}
			} else if iMid == iNumsLen-1 {
				iLeft = iMid
			}
		} else {
			iRight = iMid - 1
		}
	}
	iLast := iLeft

	//要判断没找到的情况
	if iFirst < 0 || iLast > iNumsLen-1 {
		return []int{-1, -1}
	}
	if iFirst > iLast {
		return []int{-1, -1}
	} else {
		if nums[iFirst] == target && nums[iLast] == target {
			return []int{iFirst, iLast}
		} else {
			return []int{-1, -1}
		}
	}
}
