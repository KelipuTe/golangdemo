package main

import "fmt"

func main() {
  // fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
  // fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))
  // fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
  fmt.Println(exist([][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}, "AAB"))
}

//给定一个m行n列的二维字符网格board和一个字符串单词word。如果word存在于网格中，返回true，否则，返回false
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中相邻单元格是那些水平相邻或垂直相邻的单元格
//同一个单元格内的字母不允许被重复使用
//m==board.length;n=board[i].length;1<=m,n<=6;1<=word.length<=15;board和word仅由大小写英文字母组成

//回溯
//每次回溯都有4种选择，上右下左，碰到边界的时候终止回溯
//字母不允许被重复使用，所以需要一个二维数组用于标记已经使用过的字母

var bFound bool //结果

//79-单词搜索
func exist(board [][]byte, word string) bool {
  iM, iN := len(board), len(board[0])
  var bsli2Visit [][]bool = make([][]bool, iM) //访问过的元素
  for ii := 0; ii < len(board); ii++ {
    bsli2Visit[ii] = make([]bool, iN)
  }
  bFound = false //初始化
  for ii := 0; ii < iM; ii++ {
    for ij := 0; ij < iN; ij++ {
      hui2su4(board, iM, iN, ii, ij, word, bsli2Visit)
    }
  }
  return bFound
}

func hui2su4(board [][]byte, iM int, iN int, ii int, ij int, word string, bsli2Visit [][]bool) {
  if board[ii][ij] == word[0] {
    if len(word) > 1 {
      bsli2Visit[ii][ij] = true //标记
      tsWord := word[1:]        //找剩下的字符串
      //上
      if ii > 0 && !bsli2Visit[ii-1][ij] {
        hui2su4(board, iM, iN, ii-1, ij, tsWord, bsli2Visit)
      }
      //右
      if ij < iN-1 && !bsli2Visit[ii][ij+1] {
        hui2su4(board, iM, iN, ii, ij+1, tsWord, bsli2Visit)
      }
      //下
      if ii < iM-1 && !bsli2Visit[ii+1][ij] {
        hui2su4(board, iM, iN, ii+1, ij, tsWord, bsli2Visit)
      }
      //左
      if ij > 0 && !bsli2Visit[ii][ij-1] {
        hui2su4(board, iM, iN, ii, ij-1, tsWord, bsli2Visit)
      }
      bsli2Visit[ii][ij] = false //标记复位
    } else {
      bFound = true
    }
  }
}
