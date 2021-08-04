package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{-2, 0}))
	// fmt.Println(maxProduct([]int{-2, 0, -1}))
	// fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	// fmt.Println(maxProduct([]int{5, 6, -3, 4, -3}))
}

//给定一个整数数组nums ，找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积

//152-乘积最大子数组(53,152)
func maxProduct(nums []int) int {
	//动态规划
	//可以借鉴第53题的思路，不过这里需要注意的有两点，一个是负负得正的情况，另一个是0值
	//构造两个数组，分别表示前n个元素中，连续子数组的最大乘积和最小乘积
	//初始化第一个元素，因为只有一个元素，所以连续子数组的最大乘积和最小乘积就是自己
	//从第2个元素开始，需要计算前面的最大乘积和最小乘积的结果与自身乘积，
	//找出最大乘积与自身的乘积和最小乘积与自身的乘积和自身，三个数中最大的和最小的

	var mapMaxRes, mapMinRes map[int]int = map[int]int{}, map[int]int{}
	mapMaxRes[0], mapMinRes[0] = nums[0], nums[0]
	iMaxSliMax := nums[0]
	for ii := 1; ii < len(nums); ii++ {
		if nums[ii] == 0 {
			mapMaxRes[ii], mapMinRes[ii] = 0, 0
		} else {
			tmax := mapMaxRes[ii-1] * nums[ii]
			tmin := mapMinRes[ii-1] * nums[ii]

			if tmax >= tmin && tmax >= nums[ii] {
				mapMaxRes[ii] = tmax
			}
			if tmin >= tmax && tmin >= nums[ii] {
				mapMaxRes[ii] = tmin
			}
			if nums[ii] >= tmax && nums[ii] >= tmin {
				mapMaxRes[ii] = nums[ii]
			}

			if tmax <= tmin && tmax <= nums[ii] {
				mapMinRes[ii] = tmax
			}
			if tmin <= tmax && tmin <= nums[ii] {
				mapMinRes[ii] = tmin
			}
			if nums[ii] <= tmax && nums[ii] <= tmin {
				mapMinRes[ii] = nums[ii]
			}
		}
		if mapMaxRes[ii] > iMaxSliMax {
			iMaxSliMax = mapMaxRes[ii]
		}
	}
	return iMaxSliMax
}
