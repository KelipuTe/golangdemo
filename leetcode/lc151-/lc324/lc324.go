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

//一个整数数组an，将它重新排列成an[0]<an[1]>an[2]<an[3]...的顺序

//324-摆动排序II(75,324)
func wiggleSort(nums []int) {
	//将原数组升序排序
	//将升序排序后的数组的前半部分的元素放到奇数下标，后半部分的元素放到偶数下标
	//注意，数组长度为奇数和为偶数时，下标处理方式略有差别

	iNumsLen := len(nums)
	var isli1 []int = make([]int, iNumsLen)

	//将原数组拷贝一份，然后升序排序
	copy(isli1, nums)
	sort.Ints(isli1)

	if iNumsLen%2 == 0 {
		for ii, ij := 0, (iNumsLen>>1)-1; ii < iNumsLen; {
			nums[ii] = isli1[ij]
			ii += 2
			ij--
		}
		for ii, ij := 1, iNumsLen-1; ii < iNumsLen; {
			nums[ii] = isli1[ij]
			ii += 2
			ij--
		}
	} else {
		for ii, ij := 0, iNumsLen>>1; ii < iNumsLen; {
			nums[ii] = isli1[ij]
			ii += 2
			ij--
		}
		for ii, ij := 1, iNumsLen-1; ii < iNumsLen; {
			nums[ii] = isli1[ij]
			ii += 2
			ij--
		}
	}
}
