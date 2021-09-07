//lc98-验证二叉搜索树
//[二叉树][二叉搜索树][中序遍历][数组]

//给你一个二叉树的根节点root，判断其是否是一个有效的二叉搜索树。
//有效二叉搜索树定义如下：
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。

//二叉搜索树的中序遍历是升序序列
//遍历二叉搜索树然后判断是不是升序序列

package main

import "fmt"

type TreeNode struct {
  Val   int
  Left  *TreeNode
  Right *TreeNode
}

func main() {
  phead :=
    &TreeNode{50,
      &TreeNode{30,
        &TreeNode{20, nil, nil},
        &TreeNode{40, nil, nil}},
      &TreeNode{70,
        &TreeNode{60, nil, nil},
        &TreeNode{80, nil, nil}},
    }

  fmt.Println(isValidBST(phead))
}

func isValidBST(root *TreeNode) bool {
  var sli1Nums []int = []int{}
  var f0zhong1xu4 func(root *TreeNode)

  f0zhong1xu4 = func(root *TreeNode) {
    if root == nil {
      return
    }
    if root.Left != nil {
      f0zhong1xu4(root.Left)
    }
    sli1Nums = append(sli1Nums, root.Val)
    if root.Right != nil {
      f0zhong1xu4(root.Right)
    }
  }

  f0zhong1xu4(root)

  for index := 1; index < len(sli1Nums); index++ {
    if sli1Nums[index] <= sli1Nums[index-1] {
      return false
    }
  }

  return true
}
