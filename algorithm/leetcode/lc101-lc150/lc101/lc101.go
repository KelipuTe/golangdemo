//lc101-对称二叉树
//[二叉树][递归][双指针]

//给定一个二叉树，检查它是否是镜像对称的。

//每一个结点的结果依赖左子树和右子树的结果。
//根结点不用判断，直接到第二层。
//双指针，一个往左，另一个就往右，同步遍历子树。

package main

import "fmt"

type TreeNode struct {
  Val   int
  Left  *TreeNode
  Right *TreeNode
}

func main() {
  phead := &TreeNode{1,
    &TreeNode{2,
      &TreeNode{3, nil, nil},
      &TreeNode{4, nil, nil},
    },
    &TreeNode{2,
      &TreeNode{4, nil, nil},
      &TreeNode{3, nil, nil},
    },
  }

  fmt.Println(isSymmetric(phead))
}

func isSymmetric(root *TreeNode) bool {
  var f0di4gui1 func(pleft, pright *TreeNode) bool

  f0di4gui1 = func(pleft, pright *TreeNode) bool {
    if pleft == nil && pright == nil {
      return true
    }
    if pleft == nil || pright == nil {
      return false
    }
    return pleft.Val == pright.Val && f0di4gui1(pleft.Left, pright.Right) && f0di4gui1(pleft.Right, pright.Left)
  }

  return f0di4gui1(root.Left, root.Right)
}
