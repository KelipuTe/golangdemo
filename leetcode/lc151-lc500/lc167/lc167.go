package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

//给定整数升序排列数组，元素会重复，从数组中找出两个数满足相加之和等于目标数target
//假设每个输入只对应唯一的答案，但是，数组中同一个元素不能重复出现

//167-两数之和II-输入有序数组(1,15,167)
func twoSum(numbers []int, target int) []int {
	//双指针，二分查找
	//由于数组元素会重复，所以不能用第1题的构造map的思路
	//可以参考第15题三数之和的思路，区别在于只需要在基数a1右侧找一个数a2，使得a2=n-a1就行了
	//由于数组是升序的，所以可以用二分查找找这个数

	iNumsLen := len(numbers)

	for ii := 0; ii < iNumsLen; ii++ {
		iLeft := ii + 1
		iRight := iNumsLen - 1
		for iLeft <= iRight {
			iMid := iLeft + (iRight-iLeft)>>1
			if numbers[iMid] > target-numbers[ii] {
				iRight = iMid - 1
			} else if numbers[iMid] < target-numbers[ii] {
				iLeft = iMid + 1
			} else {
				return []int{ii + 1, iMid + 1}
			}
		}
	}

	return []int{-1, -1}
}
