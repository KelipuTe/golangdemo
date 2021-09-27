//堆排序
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

  HeapSort(arr1nums, true)

  rand.Seed(time.Now().UnixNano())
  for i := 0; i < arr1numsLen; i++ {
    arr1nums[i] = rand.Intn(999)
  }

  HeapSort(arr1nums, false)
}

func HeapSort(arr1nums []int, isASC bool) {
  fmt.Println(arr1nums)

  var arr1numsLen int = len(arr1nums)
  var indexTree, indexFather int = 0, 0
  var indexEnd int = arr1numsLen //控制循环结束

  for indexEnd > 0 {
    for indexi := 1; indexi < indexEnd; indexi++ { //构造大根堆或者小根堆
      indexTree = indexi + 1 //数组下标n保存的是数组二叉树第n+1个结点
      for indexTree > 0 {
        indexFather = indexTree >> 1 //计算第n+1个结点的父结点

        if isASC { //升序排序
          if indexFather > 0 && arr1nums[indexTree-1] > arr1nums[indexFather-1] { //如果子结点比父结点大，就交换
            arr1nums[indexFather-1], arr1nums[indexTree-1] = arr1nums[indexTree-1], arr1nums[indexFather-1]
          }
        } else { //降序排序
          if indexFather > 0 && arr1nums[indexTree-1] < arr1nums[indexFather-1] { //如果子结点比父结点小，就交换
            arr1nums[indexFather-1], arr1nums[indexTree-1] = arr1nums[indexTree-1], arr1nums[indexFather-1]
          }
        }
        indexTree = indexFather //继续判断父结点
      }
    }
    arr1nums[0], arr1nums[indexEnd-1] = arr1nums[indexEnd-1], arr1nums[0] //堆化完成后，堆顶的元素是最值，把这个元素和数组末尾元素交换
    indexEnd--
  }

  fmt.Println(arr1nums)
}
