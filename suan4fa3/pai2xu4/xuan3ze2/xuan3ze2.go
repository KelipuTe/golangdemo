//选择排序
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

  SelecionSort(arr1nums, true)

  rand.Seed(time.Now().UnixNano())
  for i := 0; i < arr1numsLen; i++ {
    arr1nums[i] = rand.Intn(999)
  }

  SelecionSort(arr1nums, false)
}

func SelecionSort(arr1nums []int, isASC bool) {
  fmt.Println(arr1nums)

  var arr1numsLen int = len(arr1nums)
  for indexi := 0; indexi < arr1numsLen; indexi++ { //第n次循环找到第n小(大)的元素
    index, num := indexi, arr1nums[indexi] //本次要找的最小(大)元素的下标，最小(大)元素

    if isASC { //升序排序
      for indexj := indexi + 1; indexj < arr1numsLen; indexj++ {
        if arr1nums[indexj] < num { //如果找到更小的，就更新
          index, num = indexj, arr1nums[indexj]
        }
      }
    } else { //降序排序
      for indexj := indexi + 1; indexj < arr1numsLen; indexj++ {
        if arr1nums[indexj] > num { //如果找到更大的，就更新
          index, num = indexj, arr1nums[indexj]
        }
      }
    }
    arr1nums[index], arr1nums[indexi] = arr1nums[indexi], arr1nums[index] //把第n小(大)的元素和第n位置的元素交换
  }

  fmt.Println(arr1nums)
}
