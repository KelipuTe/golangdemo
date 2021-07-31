package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 6))
}

//704-二分查找
func search(nums []int, target int) int {
	iNumsLen := len(nums)
	iLeft, iRight := 0, iNumsLen-1
	for iLeft <= iRight {
		iMid := iLeft + (iRight-iLeft)>>1
		if nums[iMid] > target {
			iRight = iMid - 1
		} else if nums[iMid] < target {
			iLeft = iMid + 1
		} else {
			return iMid
		}
	}
	return -1
}
