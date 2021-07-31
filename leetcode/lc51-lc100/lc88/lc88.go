package main

import "fmt"

func main() {

	isli11 := []int{1, 2, 3, 0, 0, 0} //切片是引用传递，多出来的0是为了合并用的
	isli12 := []int{2, 5, 6}
	merge(isli11, 3, isli12, 3)
	fmt.Println(isli11)
}

//88-合并两个有序数组(21,88)(88,977)
func merge(nums1 []int, m int, nums2 []int, n int) {
	//使用两个下标标记两个数组中未处理的第一个数
	//每次比较两个数组中未处理的第一个数，把较小的添加到结果数组

	var tiNums []int = make([]int, 0, m+n) //结果数组
	var iIndex1, iIndex2 int = 0, 0        //下标标记

	for true {
		//其中一个数组已经到尾部
		if iIndex1 == m {
			tiNums = append(tiNums, nums2[iIndex2:]...)
			break
		}
		if iIndex2 == n {
			tiNums = append(tiNums, nums1[iIndex1:]...)
			break
		}
		//两个数组都没有到尾部
		if nums1[iIndex1] < nums2[iIndex2] {
			tiNums = append(tiNums, nums1[iIndex1])
			iIndex1++
		} else {
			tiNums = append(tiNums, nums2[iIndex2])
			iIndex2++
		}
	}

	copy(nums1, tiNums) //把结果复制到nums1中
}
