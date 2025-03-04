package main

import "fmt"

func main() {
  fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
  fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

//给定两个数组，编写一个函数来计算它们的交集。
//输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。不考虑输出结果的顺序。

//进阶：如果给定的数组已经排好序呢？你将如何优化你的算法？
//进阶：如果nums1的大小比nums2小很多，哪种方法更优？
//进阶：如果nums2的元素存储在磁盘上，内存是有限的，并且不能一次加载所有的元素到内存中，该怎么办？

//数组，哈希表
//遍历数组1，构造键为数组值，值为数组值出现次数的哈希表。
//遍历数组2，如果哈希表中存在数组值为键的元素，就将数组值入结果集，并把哈希表对应数组值出现次数-1。
//进阶：排序好的数组，可以用双指针，同时遍历两个数组。
//进阶：数组长度差距较大时，用长度小的数组构造哈希表。
//进阶：内存有限时用哈希表。

//350-两个数组的交集II
func intersect(nums1 []int, nums2 []int) []int {
  var nums1Len, nums2Len int = len(nums1), len(nums2)
  var mapNums1 map[int]int = map[int]int{}
  var sli1Res []int = []int{}

  for index := 0; index < nums1Len; index++ {
    if _, isExist := mapNums1[nums1[index]]; isExist {
      mapNums1[nums1[index]]++
    } else {
      mapNums1[nums1[index]] = 1
    }
  }

  for index := 0; index < nums2Len; index++ {
    if _, isExist := mapNums1[nums2[index]]; isExist {
      if mapNums1[nums2[index]] > 0 {
        sli1Res = append(sli1Res, nums2[index])
        mapNums1[nums2[index]]--
      }
    }
  }

  return sli1Res
}
