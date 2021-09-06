//lc226-翻转二叉树
//[二叉树][递归]

//翻转一棵二叉树。

//把每个结点的左右子树交换位置。

package main

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

  invertTree(phead)
}

func invertTree(root *TreeNode) *TreeNode {
  if root == nil {
    return root
  }
  invertTree(root.Left)
  invertTree(root.Right)
  root.Left, root.Right = root.Right, root.Left
  return root
}
