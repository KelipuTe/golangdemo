//计算平方根
package main

import (
  "fmt"
  "strconv"
)

func main() {
  var num float32 = 2  //被开方数
  var decimals int = 6 //小数位数
  var numSmall, numBig float32 = 0, num

  for times := 0; times < 20; times++ {
    var numMid float32 = numSmall + (numBig-numSmall)/2 //取中位数
    var power2 float32 = numMid * numMid

    if power2 > num { //中位数平方大于被开方数
      numBig = numMid //收缩右边界
    } else if power2 < num { //中位数平方小于被开方数
      numSmall = numMid //收缩左边界
    } else { //直接命中
      fmt.Println(numMid)
      return
    }
  }

  fmt.Println(numSmall)
  fmt.Println(numBig)

  fmt.Println(strconv.FormatFloat((float64)(numSmall+(numBig-numSmall)/2), 'f', decimals, 64))
}
