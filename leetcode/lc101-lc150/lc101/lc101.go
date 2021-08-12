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

  fmt.Println(isSymmetric(&tn1))
}

type TreeNode struct {
  Val   int
  Left  *TreeNode
  Right *TreeNode
} // 树结点

//给定一个二叉树，检查它是否是镜像对称的

//递归，双指针
//每一个结点的结果，依赖自身，左子树，右子树的结果
//根结点不用判断，直接到第二层，两个指针，一个往左，另一个就往右，同步遍历子树

//101-对称二叉树
func isSymmetric(root *TreeNode) bool {
  return di4gui1(root.Left, root.Right)
}

func di4gui1(ptnZuo, ptnYou *TreeNode) bool {
  if ptnZuo == nil && ptnYou == nil {
    return true
  }
  if ptnZuo == nil || ptnYou == nil {
    return false
  }
  return ptnZuo.Val == ptnYou.Val && di4gui1(ptnZuo.Left, ptnYou.Right) && di4gui1(ptnZuo.Right, ptnYou.Left)
}
