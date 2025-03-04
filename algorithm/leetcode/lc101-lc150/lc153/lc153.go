package main

import "fmt"

func main() {
	// fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	// fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
}

//整数数组，数组中的值互不相同，升序排列
//数组{a[0],a[1],a[2],...,a[n-1]}旋转一次的结果为数组{a[n-1],a[0],a[1],a[2],...,a[n-2]}
//例如，原数组nums=[0,1,2,4,5]，若旋转4次，则得到[4,5,0,1,2]，若旋转7次，则得到 [0,1,2,4,7]

//给你旋转后的数组nums，找出并返回数组中的最小元素
//n=nums.length;1<=n<=5000;-5000<=nums[i]<=5000

//寻找旋转排序数组中的最小值(33,81,153)
func findMin(nums []int) int {
	//二分查找
	//沿用33题的思路，二分之后，其中一半是升序的，那么最小值一定不在升序的这一半里
	//因为，升序的数组经过旋转操作后一定是中间有一个拐点的形状，而且左端点一定比右端点大
	//注意旋转n的倍数次时，旋转后的数组就是原升序数组

	iNumsLen := len(nums)

	if iNumsLen == 1 {
		return nums[0] //数组只有一个
	}
	if nums[0] <= nums[iNumsLen-1] {
		return nums[0] //数组升序
	}

	iMin := nums[0]
	il, ir := 0, iNumsLen-1
	for il <= ir {
		im := il + (ir-il)>>1
		//每次循环校验一下端点和中间点
		if nums[il] < iMin {
			iMin = nums[il]
		}
		if nums[im] < iMin {
			iMin = nums[im]
		}
		if nums[ir] < iMin {
			iMin = nums[ir]
		}
		if nums[il] <= nums[im] {
			//左半部分升序，则左半部分可以跳过
			il = im + 1
		} else {
			//右半部分升序，则右半部分可以跳过
			ir = im - 1
		}
	}
	return iMin
}
