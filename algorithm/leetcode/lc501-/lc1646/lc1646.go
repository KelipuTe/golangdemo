package main

import (
  "fmt"
)

func main() {
  fmt.Println(getMaximumGenerated(10))
}

//给一个整数n。按下述规则生成一个长度为n+1的数组nums：
//nums[0]=0;nums[1]=1;
//当2<=2*i<=n时，nums[2*i]=nums[i];
//当2<=2*i+1<=n时，nums[2*i+1]=nums[i]+nums[i+1];
//返回生成数组nums中的最大值。

//最大值一定在条件2<=2*i+1<=n这里，被构造出来

//1646-获取生成数组中的最大值
func getMaximumGenerated(n int) int {
  var sli1Res []int
  var maxNum int

  if n == 0 {
    return 0
  }

  sli1Res = make([]int, n+1)
  sli1Res[0], sli1Res[1] = 0, 1
  maxNum = 1
  for index := 0; index <= n>>1; index++ {
    if index<<1 <= n {
      sli1Res[index<<1] = sli1Res[index]
    }
    if index<<1+1 <= n {
      //最大值一定在这里被构造出来
      sli1Res[index<<1+1] = sli1Res[index] + sli1Res[index+1]
      if sli1Res[index<<1+1] > maxNum {
        maxNum = sli1Res[index<<1+1]
      }
    }
  }

  return maxNum
}
