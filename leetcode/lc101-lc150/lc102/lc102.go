package main

import "fmt"

func main() {
  tn7 := TreeNode{3, nil, nil}
  tn6 := TreeNode{4, nil, nil}
  tn5 := TreeNode{4, nil, nil}
  tn4 := TreeNode{3, nil, nil}
  tn3 := TreeNode{2, &tn6, &tn7}
  tn2 := TreeNode{2, &tn4, &tn5}
  tn1 := TreeNode{1, &tn2, &tn3}

  fmt.Println(levelOrder(&tn1))
}

type TreeNode struct {
  Val   int
  Left  *TreeNode
  Right *TreeNode
} // 树结点

//给一个二叉树，请返回其按层序遍历得到的节点值（即逐层地，从左到右访问所有节点）

//队列，广度优先遍历
//可以使用两个指针数组模拟队列操作，一个数组用于保存当前层所有的结点，另一个用于保存下一层所有的结点
//在遍历当前层的时候，记录下一层所有需要遍历的结点
//当前层遍历结束之后，用下一层的数组覆盖当前层的数组，继续遍历，直到队列为空

//102-二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
  var isli2Res [][]int = [][]int{}             //结果集
  var ptnSli1Queue []*TreeNode = []*TreeNode{} //当前层所有需要遍历的节点

  if root == nil {
    return isli2Res
  }

  ptnSli1Queue = append(ptnSli1Queue, root)
  for ii := 0; len(ptnSli1Queue) > 0; ii++ {
    isli2Res = append(isli2Res, []int{})          //结果集增加一行
    var tptnSli1Queue []*TreeNode = []*TreeNode{} //下一层所有需要遍历的节点
    for ij := 0; ij < len(ptnSli1Queue); ij++ {
      isli2Res[ii] = append(isli2Res[ii], ptnSli1Queue[ij].Val)
      if ptnSli1Queue[ij].Left != nil {
        tptnSli1Queue = append(tptnSli1Queue, ptnSli1Queue[ij].Left)
      }
      if ptnSli1Queue[ij].Right != nil {
        tptnSli1Queue = append(tptnSli1Queue, ptnSli1Queue[ij].Right)
      }
    }
    ptnSli1Queue = tptnSli1Queue
  }

  return isli2Res
}
