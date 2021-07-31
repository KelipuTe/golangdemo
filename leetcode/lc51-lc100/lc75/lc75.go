package main

import "fmt"

func main() {
	isli1 := []int{2, 1, 0, 0, 1, 2, 2, 1, 0, 2, 1}
	sortColors(isli1)
	fmt.Println(isli1)
}

//给定一个包含红色、白色和蓝色，一共n个元素的数组，使用整数0、1和2分别表示红色、白色和蓝色
//原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列
//n=nums.length;1<=n<=300;nums[i]为0,1,2

//75-颜色分类(75,324)
func sortColors(nums []int) {
	iNumsLen := len(nums)

	//从前往后遍历，找到0就换到前面
	for ii, ij := 0, 0; ij < iNumsLen; ij++ {
		if nums[ij] == 0 {
			nums[ii], nums[ij] = nums[ij], nums[ii]
			ii++
		}
	}
	//从后往前遍历，找到2就换到后面
	for ii, ij := iNumsLen-1, iNumsLen-1; ij >= 0; ij-- {
		if nums[ij] == 2 {
			nums[ii], nums[ij] = nums[ij], nums[ii]
			ii--
		}
	}
}
