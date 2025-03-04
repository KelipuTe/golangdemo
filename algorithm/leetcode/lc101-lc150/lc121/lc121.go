package main

import "fmt"

func main() {
  fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

//给定一个数组prices，它的第i个元素prices[i]表示一支给定股票第i天的价格。
//只能选择某一天买入这只股票，并选择在未来的某一个不同的日子卖出该股票。
//设计一个算法来计算你所能获取的最大利润。
//返回可以从这笔交易中获取的最大利润。如果不能获取任何利润，返回0。
//1<=prices.length<=10^5;0<=prices[i]<=10^4

//遍历数组，用当前位置的数值减去它前面所有元素中的最小值，就是当前位置的数值的最大利润

//121-买卖股票的最佳时机
func maxProfit(prices []int) int {
  var pricesLen int = len(prices)
  var max0li4run4 int = 0
  var minPrice int = prices[0]

  for index := 1; index < pricesLen; index++ {
    if prices[index]-minPrice > max0li4run4 {
      max0li4run4 = prices[index] - minPrice
    }
    if prices[index] < minPrice {
      minPrice = prices[index]
    }
  }

  return max0li4run4
}
