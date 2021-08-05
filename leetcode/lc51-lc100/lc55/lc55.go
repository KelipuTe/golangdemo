package main

import "fmt"

func main() {
	// fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	// fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{3, 2, 2, 0, 4}))
}

//给定一个非负整数数组nums ，最初位于数组的第一个下标
//数组中的每个元素代表在该位置可以跳跃的最大长度，判断是否能够到达最后一个下标
//1<=nums.length<=3*10^4;0<=nums[i]<=10^5

//55-跳跃游戏
func canJump(nums []int) bool {
	//倒推
	//从后往前找，如果后面的位置能被前面的位置推出来，就逻辑改变最后一个元素的位置
	//比如判断[2,3,1,1,4]时，4能被前面的1推出来，所以这个问题逻辑上就变成判断[2,3,1,1]
	il := len(nums) - 1
	for ii := il - 1; ii > -1; ii-- {
		if il-ii <= nums[ii] {
			il = ii
		}
	}
	return il == 0
}
