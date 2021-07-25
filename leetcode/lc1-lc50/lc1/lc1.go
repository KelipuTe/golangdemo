package main

import "fmt"

func main() {
	iarrNums := []int{2, 7, 11, 15}
	iTarget := 9
	iarrRes := twoSum(iarrNums, iTarget)
	fmt.Println(iarrRes)
}

//两数之和
//从数组an中找出两个整数[a1,a2]，使得它们的和等于给定目标数，数组元素不会重复
func twoSum(nums []int, target int) []int {
	var mapNum map[int]int = make(map[int]int, len(nums))

	//构造键为an，值为in的map
	for iIndex1, iNum1 := range nums {
		mapNum[iNum1] = iIndex1
	}

	//遍历数组，得到遍历的值a1和a2=n-a1，去map中查找有没有a2这个键
	for iNum1, iIndex1 := range mapNum {
		iNum2 := target - iNum1
		_, bExist := mapNum[iNum2]
		if bExist {
			return []int{iIndex1, mapNum[iNum2]}
		}
	}

	return []int{-1, -1}
}
