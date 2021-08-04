package main

import "fmt"

func main() {
	fmt.Println(er4fen1e([]int{2, 4, 6, 8, 10, 12, 14, 16}, 8))
}

//二分查找，查找升序数组中，等于目标数的元素的下标
func er4fen1e(nums []int, target int) int {
	iNumsLen := len(nums)
	il, ir := 0, iNumsLen-1
	for il <= ir {
		im := il + (ir-il)>>1 //取中
		if nums[im] > target {
			ir = im - 1 //目标数小于中间数，目标数在左半边，收缩右边界
		} else if nums[im] < target {
			il = im + 1 //目标数大于中间数，目标数在右半边，收缩左边界
		} else {
			return im //命中目标数
		}
	}
	return -1
}
