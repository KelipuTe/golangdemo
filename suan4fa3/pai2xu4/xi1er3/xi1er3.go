//希尔排序，又称减小增量排序，是插入排序的改进版
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

  ShellSort(arr1nums, true)

  rand.Seed(time.Now().UnixNano())
  for i := 0; i < arr1numsLen; i++ {
    arr1nums[i] = rand.Intn(999)
  }

  ShellSort(arr1nums, false)
}

func ShellSort(arr1nums []int, isASC bool) {
  fmt.Println(arr1nums)

  var arr1numsLen int = len(arr1nums)
  var stepLen int = (int)(arr1numsLen / 2) //希尔排序的分组步长
  for stepLen > 0 {
    //按分组步长进行分组，例如[0,0+n,0+2n,...]一组，每组内进行插入排序
    for indexi := stepLen; indexi < arr1numsLen; indexi++ {
      indexj := indexi
      num := arr1nums[indexj]

      //这里比较的时候，就不是和前一个元素进行比较了，而是和组内前一个元素进行比较
      if isASC {
        for indexj >= stepLen && arr1nums[indexj-stepLen] > num {
          arr1nums[indexj] = arr1nums[indexj-stepLen]
          indexj -= stepLen
        }
      } else {
        for indexj >= stepLen && arr1nums[indexj-stepLen] < num {
          arr1nums[indexj] = arr1nums[indexj-stepLen]
          indexj -= stepLen
        }
      }
      arr1nums[indexj] = num
    }
    stepLen = (int)(stepLen / 2) //重新设置一个更小的步长
  }

  fmt.Println(arr1nums)
}
