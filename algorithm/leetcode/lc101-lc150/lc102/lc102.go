//lc102-二叉树的层序遍历
//[二叉树][层序遍历][队列]

//给一个二叉树，请返回其按层序遍历得到的节点值（即逐层地，从左到右访问所有节点）

//使用队列记录每一层需要遍历的结点

package main

import "fmt"

type TreeNode struct {
  Val   int
  Left  *TreeNode
  Right *TreeNode
}

func main() {
  phead :=
    &TreeNode{1,
      &TreeNode{2,
        &TreeNode{4,
          &TreeNode{8, nil, nil},
          nil},
        &TreeNode{5, nil, nil}},
      &TreeNode{3, nil, nil},
    }

  fmt.Println(levelOrder(phead))
}

func levelOrder(root *TreeNode) [][]int {
  var sli2Res [][]int = [][]int{}
  var queue []*TreeNode = []*TreeNode{}
  var indexStart, indexEnd int = 0, 0

  if root != nil {
    queue = append(queue, root)
    indexEnd++
    ceng2Temp := -1
    for indexStart < indexEnd {
      sli2Res = append(sli2Res, []int{}) //结果集增加一行
      indexEndTemp := indexEnd           //每层遍历的结束
      ceng2Temp++
      for index := indexStart; index < indexEndTemp; index++ {
        sli2Res[ceng2Temp] = append(sli2Res[ceng2Temp], queue[index].Val)
        if queue[index].Left != nil {
          queue = append(queue, queue[index].Left)
          indexEnd++
        }
        if queue[index].Right != nil {
          queue = append(queue, queue[index].Right)
          indexEnd++
        }
      }
      indexStart = indexEndTemp //每层遍历的开始
    }
  }

  return sli2Res
}
