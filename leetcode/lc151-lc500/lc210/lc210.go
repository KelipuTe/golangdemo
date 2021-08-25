package main

import (
  "fmt"
)

func main() {
  fmt.Println(findOrder(2, [][]int{{1, 0}}))
  fmt.Println(findOrder(2, [][]int{{1, 0}, {0, 1}}))
  fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}

//现在你总共有n门课需要选，记为0到n-1。在选修某些课程之前需要一些先修课程。
//例如，想要学习课程0，你需要先完成课程1，我们用一个匹配来表示他们:[0,1]
//给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。
//可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。
//输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。
//可以假定输入的先决条件中没有重复的边。

//拓扑排序，深度优先搜索

//210-课程表II(207,210)
func findOrder(numCourses int, prerequisites [][]int) []int {
  var sli2Matrix [][]int           //有向图邻接表
  var sli1Visited []int            //顶点访问情况
  var sli1Res []int                //拓扑排序结果
  var you3huan2 bool = false       //是否有环
  var funcDFS func(indexVisit int) //深度优先搜索

  //有向图初始化并填充数据
  sli2Matrix = make([][]int, numCourses)
  for n := 0; n < len(prerequisites); n++ {
    sli2Matrix[prerequisites[n][1]] = append(sli2Matrix[prerequisites[n][1]], prerequisites[n][0])
  }

  sli1Visited = make([]int, numCourses)

  funcDFS = func(indexVisit int) {
    sli1Visited[indexVisit] = 1 //标记这个顶点正在访问
    for index := 0; index < len(sli2Matrix[indexVisit]); index++ {
      if sli1Visited[sli2Matrix[indexVisit][index]] == 0 {
        funcDFS(sli2Matrix[indexVisit][index])
        if you3huan2 {
          return
        }
      } else if sli1Visited[sli2Matrix[indexVisit][index]] == 1 {
        //如果深度遍历时，碰到被标记为正在访问的顶点，则有环
        you3huan2 = true
        return
      }
    }
    //从这个顶点出发的所有的边都访问完成，则这个顶点访问完成
    sli1Visited[indexVisit] = 2
    sli1Res = append(sli1Res, indexVisit)
  }

  //拓扑排序
  for index := 0; index < numCourses; index++ {
    if sli1Visited[index] == 0 {
      funcDFS(index)
    }
  }

  if you3huan2 {
    return []int{}
  }

  //将结果反续就是拓扑排序的结果
  for index := 0; index < numCourses>>1; index++ {
    sli1Res[index], sli1Res[numCourses-1-index] = sli1Res[numCourses-1-index], sli1Res[index]
  }

  return sli1Res
}
