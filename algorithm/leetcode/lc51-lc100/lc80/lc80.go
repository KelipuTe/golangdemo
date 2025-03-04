package main

import "fmt"

func main() {
	isli1 := []int{1, 1, 1, 2, 2, 3}
	iSli1Len := removeDuplicates(isli1)
	fmt.Println(isli1)
	fmt.Println(iSli1Len)
}

//给定一个有序数组，请原地删除重复出现的元素，使每个元素最多出现两次，返回删除后数组的新长度

//80-删除有序数组中的重复项II
func removeDuplicates(nums []int) int {
	//双指针，一个指针用于遍历，一个指针用于赋值
	//需要两个辅助参数分别记录当前记录的数组和数字已经出现的次数

	var iCheckNum int = nums[0] //校验元素
	var iNumCount int = 1       //校验计数
	var iIndexSave int = 1      //赋值位置下标

	for iIndexQuery := 1; iIndexQuery < len(nums); iIndexQuery++ {
		if iCheckNum == nums[iIndexQuery] && iNumCount < 2 {
			//校验计数小于2
			nums[iIndexSave] = nums[iIndexQuery]
			iIndexSave++
			iNumCount++
		}
		if iCheckNum != nums[iIndexQuery] {
			//校验元素变更
			nums[iIndexSave] = nums[iIndexQuery]
			iIndexSave++
			iCheckNum = nums[iIndexQuery]
			iNumCount = 1
		}
	}

	return iIndexSave
}
