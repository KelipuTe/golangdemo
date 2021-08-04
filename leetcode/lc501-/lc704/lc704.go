package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 5))
}

//704-二分查找
func search(nums []int, target int) int {
	iNumsLen := len(nums)
	il, ir := 0, iNumsLen-1
	for il <= ir {
		im := il + (ir-il)>>1
		if nums[im] > target {
			ir = im - 1
		} else if nums[im] < target {
			il = im + 1
		} else {
			return im
		}
	}
	return -1
}
