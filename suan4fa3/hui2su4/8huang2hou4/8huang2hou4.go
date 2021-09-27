//8皇后EightQueens
package main

import (
  "fmt"
  "math"
)

func main() {
  var queenNum int = 8                             //皇后数量
  var count int = 0                                //结果计数，8皇后问题，没有去重的结果，理论上有92个
  var arr1Coordinate []int = make([]int, queenNum) //[index]=value表示[index,value]位置放了一个皇后
  eightQueens(queenNum, &count, arr1Coordinate, 0)
}

func eightQueens(queenNum int, pCount *int, arr1Coordinate []int, line int) {
  for indexi := 0; indexi < queenNum; indexi++ {
    if checkCoordinate(arr1Coordinate, line, indexi) {
      arr1Coordinate[line] = indexi //line行indexi列，可以放皇后
      if line == queenNum-1 {       //已经到最后一个皇后
        (*pCount) += 1
        printMap(queenNum, pCount, arr1Coordinate) //可以输出结果
        arr1Coordinate[line] = 0                   //重置这一行
        return
      } else { // 如果还没有到最后一个皇后
        eightQueens(queenNum, pCount, arr1Coordinate, line+1) //继续求解下一行
        arr1Coordinate[line] = 0                              // 无论有没有解，这行得到求解结束，重置这一行
      }
    }
    // 继续求解这行的下一个位置，或者循环结束回到上一行继续求解
  }
}

//校验[x,y]位置能不能放皇后
func checkCoordinate(arr1Coordinate []int, x int, y int) bool {
  for indexi := 0; indexi < x; indexi++ { //第x行的皇后只需要和前面x-1行进行比较
    if arr1Coordinate[indexi] == y { //同一列，纵坐标相等
      return false
    }
    if math.Abs((float64)(indexi-x)) == math.Abs((float64)(arr1Coordinate[indexi]-y)) { //同一斜线，横坐标和纵坐标的差值的绝对值相同
      return false
    }
  }
  return true
}

//输出棋盘
func printMap(queenNum int, pCount *int, arr1Coordinate []int) {
  fmt.Printf("第%d个解：\n", *pCount)
  for indexi := 0; indexi < queenNum; indexi++ {
    for indexj := 0; indexj < queenNum; indexj++ {
      if indexj == arr1Coordinate[indexi] {
        fmt.Printf("1,")
      } else {
        fmt.Printf("0,")
      }
    }
    fmt.Printf("\n")
  }
}
