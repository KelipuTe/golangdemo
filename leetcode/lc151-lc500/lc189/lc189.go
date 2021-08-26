package main

import (
  "fmt"
)

func main() {
  var sli1 []int

  sli1 = []int{1, 2, 3, 4, 5}
  rotate(sli1, 0)
  fmt.Println(sli1)

  sli1 = []int{1, 2, 3, 4, 5}
  rotate(sli1, 3)
  fmt.Println(sli1)

  sli1 = []int{1, 2, 3, 4, 5}
  rotate(sli1, 7)
  fmt.Println(sli1)

  sli1 = []int{1}
  rotate(sli1, 0)
  fmt.Println(sli1)
}

//给定一个数组，将数组中的元素向右移动k个位置，其中k是非负数。
//进阶：至少有三种不同的方法可以解决这个问题。
//进阶：使用空间复杂度为O(1)的原地算法解决这个问题。

//数组
//举例，数组[1,2,3,4,5]，旋转2次。
//旋转2次，原来在最后面2个数，4和5，被转到了最前面2个位置。
//也就是说，旋转k次，会把最后面k个数，旋转到最前面k个位置。
//所以，把数组倒序成[5,4,3,2,1]，这样最后面的2个数字就到前面了。
//然后，把最前面2个和剩下的3（5-2）个，再倒序，就是最终的结果了。
//注意，当k大于数组长度n时，每旋转数组长度n次，数组转了一圈变回原样。
//所以，只需要处理（k%n）次就行了。

//189-旋转数组
func rotate(nums []int, k int) {
  var numsLen int = len(nums)
  //反转数组
  for index := 0; index < numsLen>>1; index++ {
    nums[index], nums[numsLen-1-index] = nums[numsLen-1-index], nums[index]
  }
  k = k % numsLen //当k大于数组长度n时，前n次操作相当于什么都没干
  //反转前k个
  for index := 0; index < k>>1; index++ {
    nums[index], nums[k-1-index] = nums[k-1-index], nums[index]
  }
  //反转剩下的
  for index := k; index < numsLen-((numsLen-k)>>1); index++ {
    nums[index], nums[numsLen-1-(index-k)] = nums[numsLen-1-(index-k)], nums[index]
  }
}
