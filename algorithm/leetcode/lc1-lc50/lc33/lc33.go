package main

import "fmt"

func main() {
	// fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	// fmt.Println(search([]int{1, 3}, 1))
	fmt.Println(search([]int{3, 1}, 3))
}

//整数数组，数组中的值互不相同，升序排列
//在传递给函数之前，数组在预先未知的某个下标k上进行了旋转
//使数组变为{an[k],an[k+1],...,an[n-1],an[0],an[1],...,an[k-1]}
//例如，[0,1,2,4,5,6,7]在下标3处经旋转后变为[4,5,6,7,0,1,2]

//给你旋转后的数组nums和一个整数target，如果nums中存在target ，则返回它的下标，否则返回-1
//1<=nums.length<=5000;-10^4<=nums[i]<=10^4;-10^4<=target<=10^4

//33-搜索旋转排序数组(33,81,153)
func search(nums []int, target int) int {
	//二分查找
	//二分之后，其中一半是升序的
	//例如，[3,4,5,6,7,0,1,2]，二分之后，左[3,4,5]和mid[6]和右[7,1,2,3]
	//这样就可以利用升序的那部分判断target在哪边

	iNumsLen := len(nums)

	if iNumsLen == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	iLeft, iRight := 0, iNumsLen-1
	for iLeft <= iRight {
		iMid := iLeft + (iRight-iLeft)>>1
		if nums[iMid] == target {
			return iMid
		}
		if nums[iLeft] <= nums[iMid] {
			//左半部分升序
			if nums[iLeft] <= target && target < nums[iMid] {
				iRight = iMid - 1
			} else {
				iLeft = iMid + 1
			}
		} else {
			//右半部分升序
			if nums[iMid] < target && target <= nums[iRight] {
				iLeft = iMid + 1
			} else {
				iRight = iMid - 1
			}
		}
	}
	return -1
}
