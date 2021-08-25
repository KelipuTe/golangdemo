package main

import (
  "fmt"
)

func main() {
  // fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
  fmt.Println(allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
}

//给你一个有n个节点的有向无环图（DAG），请你找出所有从节点0到节点n-1的路径并输出（不要求按特定顺序）
//二维数组的第i个数组中的单元都表示有向图中i号节点所能到达的下一些节点，空就是没有下一个结点了。
//n==graph.length;2<=n<=150<=graph[i][j]<n;graph[i][j]!=i（即，不存在自环）;
//graph[i]中的所有元素互不相同;保证输入为有向无环图（DAG）

//递归，深度优先搜索

//797-所有可能的路径
func allPathsSourceTarget(graph [][]int) [][]int {
  var sli2Res [][]int //路径结果集
  var graphLen int = len(graph)
  var sli1Res []int = []int{0} //深度优先搜索的每一条路径的结果

  var funcDFS func(indexVisit int, sli1Res []int) //深度优先搜索

  funcDFS = func(indexVisit int, sli1Res []int) {
    //已经遍历到目标顶点
    if indexVisit == graphLen-1 {
      var sli1ResTemp []int = make([]int, len(sli1Res))
      copy(sli1ResTemp, sli1Res)
      sli2Res = append(sli2Res, sli1ResTemp)
      return
    }
    //如果没有遍历到目标顶点，就继续递归
    for index := 0; index < len(graph[indexVisit]); index++ {
      sli1Res = append(sli1Res, graph[indexVisit][index])
      funcDFS(graph[indexVisit][index], sli1Res)
      sli1Res = sli1Res[:len(sli1Res)-1]
    }
  }

  funcDFS(0, sli1Res)

  return sli2Res
}
