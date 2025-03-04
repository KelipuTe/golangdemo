package main

import (
  "fmt"
)

func main() {
  fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1))
  fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0))
}

//动态规划，最短路径，迪杰斯特拉（Dijkstra）
//如果忽略题目中的中转次数的限制，可以把整个问题看成，在有向带权图中，寻找从某一个顶点到另外一个顶点的最短路径。
//加上中转k次的限制，可以理解成，要寻找的最短路径最多有k+1条边。
//最短路径算法中比较符合这个问题场景的是Dijkstra算法，Dijkstra算法可以计算一个顶点到其他所有顶点的最短路径。
//但是不能直接套用Dijkstra算法，借助Dijkstra算法的思想，可以使用动态规划的方式求解。
//用一个二维数组[k][n]，保存k次中转到达下标为n的城市的最小价格。
//那么中转k次到达城市n的最小价格，要么和中转k-1次一样，要么是中转k-1次到达城市i的最小价格加上城市i到城市n的价格。
//找到中转k次和中转k-1次的递推关系，就可以使用动态规划了。

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
    //用中转k-1次的结果，初始化中转k次的结果
    for index := n; index < n; index++ {
      sli2Res[times][n] = sli2Res[times-1][n]
    }
    //遍历所有的有向边
    for _, value := range flights {
      cityIdSrc, cityIdDst, price := value[0], value[1], value[2]
      //依据中转k-1次的结果，计算从cityIdSrc顶点出发到cityIdDst顶点的最小值
      if sli2Res[times-1][cityIdSrc]+price < sli2Res[times][cityIdDst] {
        sli2Res[times][cityIdDst] = sli2Res[times-1][cityIdSrc] + price
      }
    }
  }

  //遍历中转k次的结果找到最小价格
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
