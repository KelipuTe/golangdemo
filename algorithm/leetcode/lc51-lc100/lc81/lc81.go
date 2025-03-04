package main

import "fmt"

func main() {
	fmt.Println(search([]int{3, 1, 2, 3, 3, 3, 3}, 2))
}

//整数数组，数组中的值不是互不相同，非降序排列
//在传递给函数之前，数组在预先未知的某个下标k上进行了旋转
//使数组变为{an[k],an[k+1],...,an[n-1],an[0],an[1],...,an[k-1]}

//给你旋转后的数组nums和一个整数target，如果nums中存在这个目标值target，则返回true，否则返回false
//1<=nums.length<=5000;-10^4<=nums[i]<=10^4;-10^4<=target<=10^4

//81-搜索旋转排序数组II(33,81,153)
func search(nums []int, target int) bool {
	//二分查找
	//沿用33题的思路，但是与33题不同，二分之后，可能存在无法分辨哪一半是升序的情况
	//例如，nums=[3,1,2,3,3,3,3]，二分之后，左[3,1,2]和mid[3]和右[3,3,3]，无法判断
	//这个时候可以左右各自向里移动一位，变成，[1,2,3,3,3]，然后就可以使用与33题一样的处理办法了

	iNumsLen := len(nums)

	if iNumsLen == 1 {
		if nums[0] == target {
			return true
		} else {
			return false
		}
	}

	iLeft, iRight := 0, iNumsLen-1
	for iLeft <= iRight {
		iMid := iLeft + (iRight-iLeft)>>1
		if nums[iMid] == target {
			return true
		}
		if nums[iLeft] == nums[iMid] && nums[iMid] == nums[iRight] {
			//左右各自向里移动一位
			iLeft += 1
			iRight -= 1
		} else {
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
	}
	return false
}
