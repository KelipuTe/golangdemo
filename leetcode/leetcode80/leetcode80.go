package main

import "fmt"

func main() {
	isli1 := []int{1, 1, 1, 2, 2, 3}

	iLen := removeDuplicates(isli1)

	fmt.Println(isli1)
	fmt.Println(iLen)
}

//例：
//给定一个有序数组，请原地删除重复出现的元素，使每个元素最多出现两次，返回删除后数组的新长度

//解：
//双指针，一个指针用于遍历，一个指针用于赋值
//需要两个辅助参数分别记录当前记录的数组和数字已经出现的次数

//删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	iCheck := nums[0]
	iCount := 1
	ij := 1

	for ii := 1; ii < len(nums); ii++ {
		if iCheck == nums[ii] && iCount < 2 {
			nums[ij] = nums[ii]
			iCount++
			ij++
		}
		if iCheck != nums[ii] {
			nums[ij] = nums[ii]
			iCheck = nums[ii]
			iCount = 1
			ij++
		}
	}

	return ij
}
