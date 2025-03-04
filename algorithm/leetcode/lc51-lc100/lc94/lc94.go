//lc94-二叉树的中序遍历
//[二叉树][中序遍历]

//给你二叉树的根节点root，返回它节点值的中序遍历。
//树中节点数目在范围[0,100]内；-100<=Node.val<=100

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

  fmt.Println(inorderTraversal(phead))
}

func inorderTraversal(root *TreeNode) []int {
  var sli1Res []int = []int{}
  var f0zhong1xu4 func(root *TreeNode)

  f0zhong1xu4 = func(root *TreeNode) {
    if root == nil {
      return
    }
    if root.Left != nil {
      f0zhong1xu4(root.Left)
    }
    sli1Res = append(sli1Res, root.Val)
    if root.Right != nil {
      f0zhong1xu4(root.Right)
    }
  }

  f0zhong1xu4(root)

  return sli1Res
}
