package main

import "fmt"

func main() {
  fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}

type TreeNode struct {
  Val   int
  Left  *TreeNode
  Right *TreeNode
} // 树结点

//给定一棵树的前序遍历preorder与中序遍历inorder。请构造二叉树并返回其根节点

//递归
//前序遍历，根左右，中序遍历，左右根
//前序遍历的结果可以分为，根，左子树，右子树；中序遍历的结果可以分为，左子树，根，右子树
//通过根结点，可以将树的构造分为，构造左子树和构造右子树两个子问题

//105-从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
  if len(preorder) == 0 {
    return nil
  }
  //构造根结点
  tptn := &TreeNode{preorder[0], nil, nil}
  //找到中序遍历结果中，左右子树分界的位置
  ii := 0
  for ; ii < len(inorder); ii++ {
    if inorder[ii] == preorder[0] {
      break
    }
  }
  //分别把左子树和右子树的前序遍历和中序遍历结果提取出来，递归构造下面的结点
  tptn.Left = buildTree(preorder[1:len(inorder[:ii])+1], inorder[:ii])
  tptn.Right = buildTree(preorder[len(inorder[:ii])+1:], inorder[ii+1:])

  return tptn
}
