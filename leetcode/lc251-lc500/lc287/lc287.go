package main

import (
  "fmt"
)

func main() {
  fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}

//给定一个包含n+1个整数的数组nums，其数字都在1到n之间（包括1和n），可知至少存在一个重复的整数。
//假设nums只有一个重复的整数，找出这个重复的数。解决方案必须不修改数组nums且只用常量级O(1)的额外空间。
//1<=n<=105;nums.length==n+1;1<=nums[i]<=n;nums中只有一个整数出现两次或多次，其余整数均只出现一次
//进阶：如何证明nums中至少存在一个重复的数字。设计一个线性级时间复杂度O(n)的解决方案。

//图，环形链表，双指针
//根据题意，n+1个1到n之间的数字，肯定有至少一个是重复的。
//为数组每一个下标i，构造一个下标i指向下标nums[i]的边。
//例如：[1, 3, 4, 2, 2]，构造五条边为[0->1, 1->3, 2->4, 3->2, 4->2]。
//以n+1个顶点和构造出来的n+1条边构造一个图，构造出来的图里面一定存在一条带环的链表。
//这里就可以用快慢指针判断环形链表和找入环点的思路。
//根据题意，数字都在1到n之间，所以不存在下标0指向0的问题，可以从下标0位置开始切入遍历。
//从其他下标切入，存在下标i指向i的问题，例如：[3, 1, 2, 3, 4]。

//287-寻找重复数(141,142)
func findDuplicate(nums []int) int {
  var indexSlow, indexFast, indexSlow2 int

  //判断环形链表
  indexSlow, indexFast = nums[indexSlow], nums[nums[indexFast]]
  for indexSlow != indexFast {
    indexSlow, indexFast = nums[indexSlow], nums[nums[indexFast]]
  }
  //找入环点
  indexSlow2 = 0
  for indexSlow != indexSlow2 {
    indexSlow = nums[indexSlow]
    indexSlow2 = nums[indexSlow2]
  }

  return indexSlow
}
