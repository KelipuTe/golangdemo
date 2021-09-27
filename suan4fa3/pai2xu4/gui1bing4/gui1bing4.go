//归并排序
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  var arr1numsLen = 20
  var arr1nums []int = make([]int, arr1numsLen)

  rand.Seed(time.Now().UnixNano())
  for i := 0; i < arr1numsLen; i++ {
    arr1nums[i] = rand.Intn(999)
  }

  MergeSort(arr1nums, true)

  rand.Seed(time.Now().UnixNano())
  for i := 0; i < arr1numsLen; i++ {
    arr1nums[i] = rand.Intn(999)
  }

  MergeSort(arr1nums, false)
}

func MergeSort(arr1nums []int, isASC bool) {
  fmt.Println(arr1nums)
  fen1zu3(arr1nums, 0, len(arr1nums)-1, isASC)
  fmt.Println(arr1nums)
}

//分组
func fen1zu3(arr1nums []int, indexStart int, indexEnd int, isASC bool) {
  if indexStart < indexEnd {
    var indexMid int = indexStart + (indexEnd-indexStart)>>1
    fen1zu3(arr1nums, indexStart, indexMid, isASC)
    fen1zu3(arr1nums, indexMid+1, indexEnd, isASC)
    he2bing4(arr1nums, indexStart, indexEnd, isASC)
  }
}

//合并
func he2bing4(arr1nums []int, indexStart int, indexEnd int, isASC bool) {
  var arr1numsLen int = len(arr1nums)
  var arr1numsTemp []int = make([]int, arr1numsLen) //临时保存本次归并的结果
  var indexTemp int = indexStart
  var indexMid int = indexStart + (indexEnd-indexStart)>>1
  var indexi, indexj int = indexStart, indexMid + 1

  for indexi < indexMid+1 && indexj < indexEnd+1 { //两半序列都没有遍历结束时，每次取出头部的元素进行比较
    if isASC { //升序排序
      if arr1nums[indexi] > arr1nums[indexj] {
        arr1numsTemp[indexTemp] = arr1nums[indexj]
        indexj++
      } else {
        arr1numsTemp[indexTemp] = arr1nums[indexi]
        indexi++
      }
    } else { //降序排序
      if arr1nums[indexi] > arr1nums[indexj] {
        arr1numsTemp[indexTemp] = arr1nums[indexi]
        indexi++
      } else {
        arr1numsTemp[indexTemp] = arr1nums[indexj]
        indexj++
      }
    }
    indexTemp++
  }

  //下面这两个循环每次只会走一个
  for indexi < indexMid+1 { //后半部分序列遍历结束，直接添加前半序列的剩余元素
    arr1numsTemp[indexTemp] = arr1nums[indexi]
    indexi++
    indexTemp++
  }
  for indexj < indexEnd+1 { //前半部分序列遍历结束，直接添加前半序列的剩余元素
    arr1numsTemp[indexTemp] = arr1nums[indexj]
    indexj++
    indexTemp++
  }

  //把排序好的序列替换到原序列的位置
  for index := indexStart; index < indexEnd+1; index++ {
    arr1nums[index] = arr1numsTemp[index]
  }
}
