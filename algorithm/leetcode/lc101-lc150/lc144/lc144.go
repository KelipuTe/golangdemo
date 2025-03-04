//lc144-二叉树的前序遍历
//[二叉树][前序遍历]

//给你二叉树的根节点root，返回它节点值的前序遍历。
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

  fmt.Println(preorderTraversal(phead))
}

func preorderTraversal(root *TreeNode) []int {
  var sli1Res []int = []int{}
  var f0qian2xu4 func(root *TreeNode)

  f0qian2xu4 = func(root *TreeNode) {
    if root == nil {
      return
    }
    sli1Res = append(sli1Res, root.Val)
    if root.Left != nil {
      f0qian2xu4(root.Left)
    }
    if root.Right != nil {
      f0qian2xu4(root.Right)
    }
  }

  f0qian2xu4(root)

  return sli1Res
}
