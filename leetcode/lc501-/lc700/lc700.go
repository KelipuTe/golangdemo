//lc700-二叉搜索树中的搜索
//[二叉树][二叉搜索树][递归]

//给定二叉搜索树（BST）的根节点和一个值。需要在BST中找到节点值等于给定值的节点。
//返回以该节点为根的子树。如果节点不存在，则返回NULL。

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

  fmt.Println(searchBST(phead, 20))
  fmt.Println(searchBST(phead, 30))
  fmt.Println(searchBST(phead, 60))
  fmt.Println(searchBST(phead, 65))
  fmt.Println(searchBST(phead, 25))
}

func searchBST(root *TreeNode, val int) *TreeNode {
  if root == nil {
    return nil
  }
  if val < root.Val {
    if root.Left != nil {
      return searchBST(root.Left, val)
    }
  } else if val > root.Val {
    if root.Right != nil {
      return searchBST(root.Right, val)
    }
  } else {
    return root
  }
  return nil
}
