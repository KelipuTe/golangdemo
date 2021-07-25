package main

import (
	"fmt"
	"sort"
)

func main() {
	isli1 := []int{1, 5, 1, 1, 6, 4}
	// isli1 := []int{1, 3, 2, 2, 3, 1}
	// isli1 := []int{1, 1, 2, 1, 2, 2, 1}
	// isli1 := []int{4, 5, 5, 6}
	wiggleSort(isli1)
	fmt.Println(isli1)
}

//摆动排序
//一个整数数组an，将它重新排列成 an[0] < an[1] > an[2] < an[3] ... 的顺序
func wiggleSort(nums []int) {
	iSliLen := len(nums)
	var isli1 []int = make([]int, iSliLen)

	//将原数组拷贝一份，然后升序排序
	copy(isli1, nums)
	sort.Ints(isli1)

	//将前半部分的元素放到奇数下标，后半部分的元素放到偶数下标
	//注意，数组长度为奇数和为偶数时，下标处理方式略有差别
	if iSliLen%2 == 0 {
		for ii, ij := 0, iSliLen/2-1; ii < iSliLen; {
			nums[ii] = isli1[ij]
			ii += 2
			ij--
		}
		for ii, ij := 1, iSliLen-1; ii < iSliLen; {
			nums[ii] = isli1[ij]
			ii += 2
			ij--
		}
	} else {
		for ii, ij := 0, iSliLen/2; ii < iSliLen; {
			nums[ii] = isli1[ij]
			ii += 2
			ij--
		}
		for ii, ij := 1, iSliLen-1; ii < iSliLen; {
			nums[ii] = isli1[ij]
			ii += 2
			ij--
		}
	}
}
