//lc701-二叉搜索树中的插入操作
//[二叉树][二叉搜索树][递归]

//给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。
//返回插入后二叉搜索树的根节点。输入数据保证，新值和原始二叉搜索树中的任意节点值都不同。
//注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。
//可以返回任意有效的结果。

package main

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

  insertIntoBST(phead, 20)
  insertIntoBST(phead, 25)
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
  if root == nil {
    return &TreeNode{val, nil, nil}
  }
  if val < root.Val {
    if root.Left != nil {
      insertIntoBST(root.Left, val)
    } else {
      root.Left = &TreeNode{val, nil, nil}
    }
  } else if val > root.Val {
    if root.Right != nil {
      insertIntoBST(root.Right, val)
    } else {
      root.Right = &TreeNode{val, nil, nil}
    }
  }
  return root
}
