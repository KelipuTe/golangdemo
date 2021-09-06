//lc112-路径总和
//[二叉树][递归]

//给二叉树的根节点root和一个表示目标和的整数targetSum，
//判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和targetSum。
//叶子节点是指没有子节点的节点。

//递归分解成子问题

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

  fmt.Println(hasPathSum(nil, 8))
  fmt.Println(hasPathSum(phead, 8))
  fmt.Println(hasPathSum(phead, 9))
  fmt.Println(hasPathSum(phead, 4))

  // phead :=
  //   &TreeNode{-2,
  //     nil,
  //     &TreeNode{-3, nil, nil},
  //   }

  // fmt.Println(hasPathSum(phead, -5))
}

func hasPathSum(root *TreeNode, targetSum int) bool {
  if root == nil {
    return false
  }
  if root.Val == targetSum && root.Left == nil && root.Right == nil {
    return true
  }
  if root.Left != nil {
    if hasPathSum(root.Left, targetSum-root.Val) {
      return true
    }
  }
  if root.Right != nil {
    if hasPathSum(root.Right, targetSum-root.Val) {
      return true
    }
  }
  return false
}
