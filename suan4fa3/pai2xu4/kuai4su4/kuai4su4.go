//快速排序
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  var arr1numsLen = 10
  var arr1nums []int = make([]int, arr1numsLen)

  rand.Seed(time.Now().UnixNano())
  for i := 0; i < arr1numsLen; i++ {
    arr1nums[i] = rand.Intn(999)
  }

  fmt.Println(arr1nums)
  fmt.Println(QuickSortMinK(arr1nums, 0, len(arr1nums)-1, 4))
  QuickSort(arr1nums, 0, len(arr1nums)-1, true)
  fmt.Println(arr1nums)

  rand.Seed(time.Now().UnixNano())
  for i := 0; i < arr1numsLen; i++ {
    arr1nums[i] = rand.Intn(999)
  }

  fmt.Println(arr1nums)
  QuickSort(arr1nums, 0, len(arr1nums)-1, false)
  fmt.Println(arr1nums)
}

func QuickSort(arr1nums []int, indexStart int, indexEnd int, isASC bool) {
  if indexStart >= indexEnd {
    return
  }

  var midNum = arr1nums[indexStart] //选第一个元素做为基准元素
  var indexi, indexj int = indexStart, indexEnd

  for indexi < indexj { //从后往前找一个比基准元素小的，交换位置
    if isASC { //升序排序
      if arr1nums[indexj] < midNum {
        arr1nums[indexi], arr1nums[indexj] = arr1nums[indexj], arr1nums[indexi]
        indexi++
        for indexi < indexj { //从前往后找一个比基准元素大的，交换位置
          if arr1nums[indexi] > midNum {
            arr1nums[indexi], arr1nums[indexj] = arr1nums[indexj], arr1nums[indexi]
            break //只要找到第一个就结束本轮循环
          }
          indexi++
        }
      }
      indexj--
      //循环结束后，序列分成了三部分，比基准元素小的，基准元素，比基准元素大的
    } else { //降序排序
      if arr1nums[indexj] > midNum {
        arr1nums[indexi], arr1nums[indexj] = arr1nums[indexj], arr1nums[indexi]
        indexi++
        for indexi < indexj { //从前往后找一个比基准元素大的，交换位置
          if arr1nums[indexi] < midNum {
            arr1nums[indexi], arr1nums[indexj] = arr1nums[indexj], arr1nums[indexi]
            break //只要找到第一个就结束本轮循环
          }
          indexi++
        }
      }
      indexj--
      //循环结束后，序列分成了三部分，比基准元素大的，基准元素，比基准元素小的
    }
  }

  QuickSort(arr1nums, indexStart, indexi, isASC)
  QuickSort(arr1nums, indexStart+1, indexEnd, isASC)
}

//第k小的元素（第k大的元素同理）
func QuickSortMinK(arr1nums []int, indexStart int, indexEnd int, k int) int {
  if indexStart >= indexEnd {
    return -1
  }

  var midNum = arr1nums[indexStart]
  var indexi, indexj int = indexStart, indexEnd

  for indexi < indexj {
    if arr1nums[indexj] < midNum {
      arr1nums[indexi], arr1nums[indexj] = arr1nums[indexj], arr1nums[indexi]
      indexi++
      for indexi < indexj {
        if arr1nums[indexi] > midNum {
          arr1nums[indexi], arr1nums[indexj] = arr1nums[indexj], arr1nums[indexi]
          break
        }
        indexi++
      }
    }
    indexj--
  }

  //前半部分和快排一样，不一样的地方是最后这里，只需要处理一部分数据
  //比较一下，确定k在哪个区间内，另外一边就不需要处理了
  if indexi > k-1 {
    return QuickSortMinK(arr1nums, indexStart, indexi, k)
  } else if indexi < k-1 {
    return QuickSortMinK(arr1nums, indexi+1, indexEnd, k)
  } else {
    return arr1nums[indexi]
  }
}
