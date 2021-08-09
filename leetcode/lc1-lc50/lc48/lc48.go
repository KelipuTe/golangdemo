package main

import "fmt"

//给定一个n×n的二维矩阵matrix表示一个图像。将图像顺时针旋转90度
//必须在原地旋转图像，不要使用另一个矩阵来旋转图像

func main() {
  isli2Matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
  rotate(isli2Matrix)
  fmt.Println(isli2Matrix)
}

//48-旋转图像
func rotate(matrix [][]int) {
  //连接矩阵的对角线将矩阵分成四个部分，这四个部分都顺时针旋转90度
  //那么旋转时，对应位置的坐标可以表示成，[i,j],[j,n-1-i],[n-1-i,n-1-j],[n-1-j,i]
  in := len(matrix)
  for ii := 0; ii < in-1; ii++ {
    for ij := ii; ij < in-1-ii; ij++ {
      matrix[ii][ij], matrix[ij][in-1-ii], matrix[in-1-ii][in-1-ij], matrix[in-1-ij][ii] =
        matrix[in-1-ij][ii], matrix[ii][ij], matrix[ij][in-1-ii], matrix[in-1-ii][in-1-ij]
    }
  }
}
