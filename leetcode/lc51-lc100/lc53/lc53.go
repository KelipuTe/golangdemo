package main

import "fmt"

func main() {
	// fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}

//给定一个整数数组nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和

//53-最大子序和(53,152)
func maxSubArray(nums []int) int {
	//动态规划
	//构造一个数组，表示前n个元素中，连续子数组的最大和
	//初始化第一个元素，因为只有一个元素，所以连续子数组最大和就是自己
	//从第2个元素开始，前面的结果与自身的和，有两种情况，比自身大或者没有自身大
	//比自身大就，把和作为前n个元素连续子数组的最大和，如果没有自身大，那么自己就是连续子数组的最大和

	var mapRes map[int]int = map[int]int{}
	mapRes[0] = nums[0]
	iMaxSliMax := nums[0]
	for ii := 1; ii < len(nums); ii++ {
		if mapRes[ii-1]+nums[ii] <= nums[ii] {
			mapRes[ii] = nums[ii]
		} else {
			mapRes[ii] = mapRes[ii-1] + nums[ii]
		}
		if mapRes[ii] > iMaxSliMax {
			iMaxSliMax = mapRes[ii]
		}
	}
	return iMaxSliMax
}
