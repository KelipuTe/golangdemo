package main

import "fmt"

func main() {
  // fmt.Println(floodFill([][]int{
  //   {1, 1, 1},
  //   {1, 1, 0},
  //   {1, 0, 1},
  // }, 1, 1, 2))

  fmt.Println(floodFill([][]int{
    {0, 0, 0},
    {0, 1, 1},
  }, 1, 1, 1))
}

//有一幅以二维整数数组表示的图画，每一个整数表示该图画的像素值大小，数值在0到65535之间。
//给你一个坐标(sr,sc)表示图像渲染开始的像素值（行，列）和一个新的颜色值newColor，让你重新上色这幅图像。
//为了完成上色工作，从初始坐标开始，记录初始坐标的上下左右四个方向上像素值与初始坐标相同的相连像素点，
//接着再记录这四个方向上符合条件的像素点与他们对应四个方向上像素值与初始坐标相同的相连像素点，
//重复该过程，将所有有记录的像素点的颜色值改为新的颜色值。最后返回经过上色渲染后的图像。

//数组，广度优先搜索，队列
//类似广度优先搜索，从出发点开始，修改，然后入队。
//如果队列不为空，就出队，在访问位置，检查上右下左四个方向。
//如果需要修改，就修改，然后入队。直到队列为空，搜索结束。
//注意，不需要修改颜色的情况。

//733-图像渲染(200,695,733)
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
  var hang2, lie4 int = len(image), len(image[0]) //二维数组行列数
  var oldColor int = image[sr][sc]                //原颜色值
  var queue [][2]int                              //队列
  var queueStart, queueEnd int = 0, 0             //队列下标

  if oldColor == newColor {
    return image
  }

  //类似广度优先搜索
  image[sr][sc] = newColor //修改颜色
  queue = append(queue, [2]int{sr, sc})
  queueEnd++
  for queueStart < queueEnd {
    hang2Temp, lie4Temp := queue[queueStart][0], queue[queueStart][1]
    queueStart++
    //检查上右下左四个方向
    if hang2Temp-1 >= 0 && image[hang2Temp-1][lie4Temp] == oldColor {
      image[hang2Temp-1][lie4Temp] = newColor
      queue = append(queue, [2]int{hang2Temp - 1, lie4Temp})
      queueEnd++
    }
    if lie4Temp+1 < lie4 && image[hang2Temp][lie4Temp+1] == oldColor {
      image[hang2Temp][lie4Temp+1] = newColor
      queue = append(queue, [2]int{hang2Temp, lie4Temp + 1})
      queueEnd++
    }
    if hang2Temp+1 < hang2 && image[hang2Temp+1][lie4Temp] == oldColor {
      image[hang2Temp+1][lie4Temp] = newColor
      queue = append(queue, [2]int{hang2Temp + 1, lie4Temp})
      queueEnd++
    }
    if lie4Temp-1 >= 0 && image[hang2Temp][lie4Temp-1] == oldColor {
      image[hang2Temp][lie4Temp-1] = newColor
      queue = append(queue, [2]int{hang2Temp, lie4Temp - 1})
      queueEnd++
    }
  }

  return image
}
