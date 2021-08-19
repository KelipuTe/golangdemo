package main

import (
	"fmt"
)

func main() {
  fmt.Println(canFinish(2, [][]int{{1, 0}}))
  fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
  fmt.Println(canFinish(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}

//你这个学期必须选修numCourses门课程，记为0到numCourses-1。
//在选修某些课程之前需要一些先修课程。先修课程按数组prerequisites给出，
//其中prerequisites[i]=[ai,bi]，表示如果要学习课程ai则必须先学习课程bi。
//例如，先修课程对[0,1]表示：想要学习课程0，你需要先完成课程1。
//请你判断是否可能完成所有课程的学习？如果可以，返回true；否则，返回false。
//1<=numCourses<=10^5;0<=prerequisites.length<=5000;prerequisites[i].length==2;
//0<=ai,bi<numCourses;prerequisites[i]中的所有课程对互不相同

//拓扑排序，深度优先搜索

//207-课程表
func canFinish(numCourses int, prerequisites [][]int) bool {
  var sli2Matrix [][]int           //伪有向图邻接表
  var sli1Visited []int            //顶点访问情况
  var you3huan2 bool = false       //是否有环
  var funcDFS func(indexVisit int) //深度优先搜索

  //有向图矩阵初始化并填充数据
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
  }

  //拓扑排序
  for index := 0; index < numCourses; index++ {
    if sli1Visited[index] == 0 {
      funcDFS(index)
    }
  }

  return !you3huan2
}
