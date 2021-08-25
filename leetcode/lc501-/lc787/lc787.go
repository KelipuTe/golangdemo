package main

import (
  "fmt"
)

func main() {
  fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1))
}

//动态规划

//787-K站中转内最便宜的航班
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
  var sli2Res [][]int     //结果集，[中转次数][到达城市]最小价格
  var maxPrice = 10000000 //临界值

  //初始化结果集
  sli2Res = make([][]int, k+2)
  for times := 0; times < k+2; times++ {
    sli2Res[times] = make([]int, n)

    for cityIdDst := 0; cityIdDst < n; cityIdDst++ {
      sli2Res[times][cityIdDst] = maxPrice //默认src顶点到其他顶点的最小价格为临界值
      sli2Res[times][src] = 0              //自己到自己初始化0
    }
  }

  //计算中转k次的结果
  for times := 1; times <= k+1; times++ {
    //遍历所有的有向边
    for _, value := range flights {
      cityIdSrc, cityIdDst, price := value[0], value[1], value[2]
      if sli2Res[times-1][cityIdSrc]+price < sli2Res[times][cityIdDst] {
        sli2Res[times][cityIdDst] = sli2Res[times-1][cityIdSrc] + price
      } else {
        sli2Res[times][cityIdDst] = sli2Res[times-1][cityIdDst]
      }
    }
  }

  var minPrice int = maxPrice
  for times := 0; times <= k+1; times++ {
    if sli2Res[times][dst] < minPrice {
      minPrice = sli2Res[times][dst]
    }
  }
  if minPrice == maxPrice {
    minPrice = -1
  }

  return minPrice
}
