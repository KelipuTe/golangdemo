//lc104-二叉树的最大深度
//[二叉树][递归]

//给定一个二叉树，找出其最大深度。
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//说明: 叶子节点是指没有子节点的节点。

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

  fmt.Println(maxDepth(phead))
}

func maxDepth(root *TreeNode) int {
  if root == nil {
    return 0
  }
  depthLeft := maxDepth(root.Left) + 1
  depthRight := maxDepth(root.Right) + 1
  if depthLeft > depthRight {
    return depthLeft
  }
  return depthRight
}
