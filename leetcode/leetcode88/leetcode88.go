package main

import "fmt"

//链表结点
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//函数是引用传递，多出来的0是为了合并用的
	isli11 := []int{1, 2, 3, 0, 0, 0}
	isli12 := []int{2, 5, 6}

	merge(isli11, 3, isli12, 3)

	for ii := 0; ii < len(isli11); ii++ {
		fmt.Printf("%d,", isli11[ii])
	}
}

//解：
//使用两个下标标记两个数组中未处理的第一个数
//每次比较两个数组中未处理的第一个数，把较小的添加到结果数组中
//下标指向下一个未处理的数

//合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	tiNums := make([]int, 0, m+n)
	ii1, ii2 := 0, 0

	for true {
		//其中一个数组已经到尾部
		if ii1 == m {
			tiNums = append(tiNums, nums2[ii2:]...)
			break
		}
		if ii2 == n {
			tiNums = append(tiNums, nums1[ii1:]...)
			break
		}
		//两个数组都没有到尾部
		if nums1[ii1] < nums2[ii2] {
			tiNums = append(tiNums, nums1[ii1])
			ii1++
		} else {
			tiNums = append(tiNums, nums2[ii2])
			ii2++
		}
	}
	copy(nums1, tiNums)
}
