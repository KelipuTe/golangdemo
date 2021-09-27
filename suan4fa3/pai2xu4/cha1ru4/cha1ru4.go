//插入排序
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

  InsertionSort(arr1nums, true)

  rand.Seed(time.Now().UnixNano())
  for i := 0; i < arr1numsLen; i++ {
    arr1nums[i] = rand.Intn(999)
  }

  InsertionSort(arr1nums, false)
}

func InsertionSort(arr1nums []int, isASC bool) {
  fmt.Println(arr1nums)

  var arr1numsLen int = len(arr1nums)
  for indexi := 1; indexi < arr1numsLen; indexi++ { //假定第一个元素有序，从第二个元素开始
    indexj, num := indexi, arr1nums[indexi] //本次参与排序的位置和元素值

    if isASC { //升序排序
      for indexj > 0 && arr1nums[indexj-1] > num { //前面的数字大
        arr1nums[indexj] = arr1nums[indexj-1] //把大数字往后移
        indexj--
      }
    } else { //降序排序
      for indexj > 0 && arr1nums[indexj-1] < num { //前面的数字小
        arr1nums[indexj] = arr1nums[indexj-1] //把小数字往后移
        indexj--
      }
    }
    arr1nums[indexj] = num
  }

  fmt.Println(arr1nums)
}
