package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(triangleNumber([]int{0, 0, 2, 2, 3, 4, 5}))
	// fmt.Println(triangleNumber([]int{1, 1, 3, 4}))
	fmt.Println(triangleNumber([]int{82, 15, 23, 82, 67, 0, 3, 92, 11}))
}

//给定一个包含非负整数的数组，统计其中可以组成三角形三条边的三元组个数
//数组长度不超过1000，数组里整数的范围为[0,1000]

//611-有效三角形的个数
func triangleNumber(nums []int) int {
	//双指针，二分查找
	//构成三角形的条件是两边之和大于第三边，设三边分别为a,b,c，需要同时满足a+b>c;a+c>b;b+c>a
	//如果将数组升序排序，假设a<=b<=c，则a+c>b;b+c>a恒成立，这时只需要考虑a+b>c这个条件
	//可以用双指针的思路，固定一条边然后遍历另外两条边，依次判断。也可以结合二分查找，优化寻找第三条边的速度

	iNumsLen := len(nums)
	iRes := 0

	sort.Ints(nums)

	for ia := 0; ia < iNumsLen-2; ia++ {
		if nums[ia] <= 0 {
			continue //非负整数数组，要跳过0
		}
		for ib := ia + 1; ib < iNumsLen-1; ib++ {
			//要求a+b>c，所以要找小于a+b的最大值的下标，可以用二分查找
			iIndex := iNumsLen
			il, ir := 0, iNumsLen-1
			for il <= ir {
				im := il + (ir-il)>>1
				if nums[im] < nums[ia]+nums[ib] {
					il = im + 1
					iIndex = im
				} else {
					ir = im - 1
				}
			}
			if iIndex > ib {
				iRes += iIndex - ib
			}
		}
	}

	return iRes
}
