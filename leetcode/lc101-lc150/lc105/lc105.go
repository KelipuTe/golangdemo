package main

import "fmt"

// 105. 从前序与中序遍历序列构造二叉树
// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

// 解题思路：递归
// 前序遍历，根左右；中序遍历，左右根
// 前序遍历的结果可以分为，根，左子树，右子树；中序遍历的结果可以分为，左子树，根，右子树
// 通过根结点，可以将树的构造分为，构造左子树和构造右子树两个子问题

func main() {
  fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
  Val   int
  Left  *TreeNode
  Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
  if len(preorder) == 0 {
    return nil
  }
  // 构造根结点
  t0p1root := &TreeNode{preorder[0], nil, nil}
  // 找到中序遍历结果中，左右子树分界的位置
  indexI := 0
  for ; indexI < len(inorder); indexI++ {
    if inorder[indexI] == preorder[0] {
      break
    }
  }
  // 分别把左子树和右子树的前序遍历和中序遍历结果提取出来，递归构造下面的结点
  t0p1root.Left = buildTree(preorder[1:len(inorder[:indexI])+1], inorder[:indexI])
  t0p1root.Right = buildTree(preorder[len(inorder[:indexI])+1:], inorder[indexI+1:])

  return t0p1root
}
