package main

import "fmt"

func main() {
	fmt.Println(er4fen1gt([]int{2, 4, 6, 8, 10, 12, 14, 16}, 10))
	fmt.Println(er4fen1gt([]int{2, 2, 4, 4, 4, 4, 4, 8, 8}, 4))
}

//二分查找，查找升序数组中，大于目标数的最小的元素的下标
func er4fen1gt(nums []int, target int) int {
	iNumsLen := len(nums)
	iIndex := iNumsLen //结果位置初始化为数组长度，如果没有变过，说明目标数比数组中任意一个元素都大
	il, ir := 0, iNumsLen-1
	for il <= ir {
		im := il + (ir-il)>>1 //取中
		if nums[im] > target {
			ir = im - 1 //目标数小于中间数，目标数在左半边，收缩右边界
			iIndex = im //需要收缩右边界，则说明大于目标数的最小的元素在中间数位置或者中间数左侧位置
		} else {
			il = im + 1 //目标数大于中间数，目标数在右半边，收缩左边界
		}
	}
	return iIndex
}
