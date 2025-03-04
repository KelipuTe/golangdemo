package main

import "fmt"

func main() {
  fmt.Println(maxAreaOfIsland([][]int{
    {0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
    {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
    {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
    {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
    {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
    {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
    {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
    {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
  }))
}

//给定一个包含了一些0和1的非空二维数组grid。
//一个岛屿是由一些相邻的1(代表土地)构成的组合，这里的相邻要求两个1必须在水平或者竖直方向上相邻。
//可以假设grid的四个边缘都被0（代表水）包围着。
//找到给定的二维数组中最大的岛屿面积。如果没有岛屿，则返回面积为0。
//给定的矩阵grid的长度和宽度都不超过50。

//数组，广度优先搜索，队列
//类似广度优先搜索，遍历二维数组，如果找到陆地，标记为访问过，然后入队。
//如果队列不为空，就出队，在访问位置，检查上右下左四个方向。
//如果是陆地，标记为访问过，然后入队。直到访问队列为空，本次搜索结束，比较最大面积。
//继续遍历二维数组，直到所有的位置都访问过一次。

//695-岛屿的最大面积(200,695,733)
func maxAreaOfIsland(grid [][]int) int {
  var hang2, lie4 int = len(grid), len(grid[0]) //二维数组行列数
  var land, visited int = 1, 2                  //陆地，访问过的陆地
  var queue [][2]int                            //队列
  var queueStart, queueEnd int = 0, 0           //队列下标
  var areaMax, areaTemp int = 0, 0

  for hang2Index := 0; hang2Index < hang2; hang2Index++ {
    for lie4Index := 0; lie4Index < lie4; lie4Index++ {
      if grid[hang2Index][lie4Index] == land {
        //类似广度优先搜索
        areaTemp = 1
        grid[hang2Index][lie4Index] = visited //标记陆地已访问
        queue = append(queue, [2]int{hang2Index, lie4Index})
        queueEnd++
        for queueStart < queueEnd {
          hang2Temp, lie4Temp := queue[queueStart][0], queue[queueStart][1]
          queueStart++
          //检查上右下左四个方向
          if hang2Temp-1 >= 0 && grid[hang2Temp-1][lie4Temp] == land {
            areaTemp++
            grid[hang2Temp-1][lie4Temp] = visited
            queue = append(queue, [2]int{hang2Temp - 1, lie4Temp})
            queueEnd++
          }
          if lie4Temp+1 < lie4 && grid[hang2Temp][lie4Temp+1] == land {
            areaTemp++
            grid[hang2Temp][lie4Temp+1] = visited
            queue = append(queue, [2]int{hang2Temp, lie4Temp + 1})
            queueEnd++
          }
          if hang2Temp+1 < hang2 && grid[hang2Temp+1][lie4Temp] == land {
            areaTemp++
            grid[hang2Temp+1][lie4Temp] = visited
            queue = append(queue, [2]int{hang2Temp + 1, lie4Temp})
            queueEnd++
          }
          if lie4Temp-1 >= 0 && grid[hang2Temp][lie4Temp-1] == land {
            areaTemp++
            grid[hang2Temp][lie4Temp-1] = visited
            queue = append(queue, [2]int{hang2Temp, lie4Temp - 1})
            queueEnd++
          }
        }
        //相连的陆地已经全部标记，比较最大面积
        if areaTemp > areaMax {
          areaMax = areaTemp
        }
      }
    }
  }

  return areaMax
}
