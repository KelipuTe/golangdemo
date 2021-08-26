package main

import "fmt"

func main() {
  sli11 := []int{1, 2, 3, 0, 0, 0} //切片是引用传递，多出来的0是为了合并用的
  sli12 := []int{2, 5, 6}
  merge(sli11, 3, sli12, 3)
  fmt.Println(sli11)
}

//给你两个按非递减顺序排列的整数数组nums1和nums2，另有两个整数m和n，分别表示nums1和nums2中的元素数目。
//请你合并nums2到nums1中，使合并后的数组同样按非递减顺序排列。
//注意：最终，合并后数组不应由函数返回，而是存储在数组nums1中。
//为了应对这种情况，nums1的初始长度为m+n，其中前m个元素表示应合并的元素，后n个元素为0，应忽略。nums2的长度为n。
//nums1.length==m+n;nums2.length==n;0<=m,n<=200;1<=m+n<=200;-10^9<=nums1[i],nums2[j]<=10^9
//进阶：实现一个时间复杂度为O(m+n)的算法解决此问题

//数组，双指针
//使用两个指针，指向两个数组中未处理的第一个数。
//每次比较两个数组中未处理的第一个数，把较小的添加到结果数组。

//88-合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
  var sli1Res []int = make([]int, 0, m+n)
  var index1, index2 int = 0, 0

  for true {
    //其中一个数组已经到尾部
    if index1 == m {
      sli1Res = append(sli1Res, nums2[index2:]...)
      break
    }
    if index2 == n {
      sli1Res = append(sli1Res, nums1[index1:]...)
      break
    }
    //两个数组都没有到尾部
    if nums1[index1] < nums2[index2] {
      sli1Res = append(sli1Res, nums1[index1])
      index1++
    } else {
      sli1Res = append(sli1Res, nums2[index2])
      index2++
    }
  }

  copy(nums1, sli1Res) //把结果复制到nums1中
}
