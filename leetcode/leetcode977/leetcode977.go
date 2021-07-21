package main

import "fmt"

func main() {
	isli1Nums := []int{-4, -1, 0, 3, 10}

	isli1Res := sortedSquares(isli1Nums)

	for ii := 0; ii < len(isli1Res); ii++ {
		fmt.Printf("%d,", isli1Res[ii])
	}
}

//解：
//负数越小平方越大，正数越大平方越大
//因为数组是有序的，所以最小的负数和最大的正数分别位于两侧
//从两侧开始，将数字平方后比较，大的那个就是结果数组里大的那个

//有序数组的平方
func sortedSquares(nums []int) []int {
	iArrLen := len(nums)
	isli1Nums := make([]int, iArrLen)
	ii, ij, ik := 0, iArrLen-1, iArrLen-1

	for ii < ij {
		if nums[ii]*nums[ii] > nums[ij]*nums[ij] {
			isli1Nums[ik] = nums[ii] * nums[ii]
			ik--
			ii++
		} else {
			isli1Nums[ik] = nums[ij] * nums[ij]
			ik--
			ij--
		}
	}

	return isli1Nums
}
