//lc235-二叉搜索树的最近公共祖先
//[二叉树][二叉搜索树][递归]

//给定一个二叉搜索树,找到该树中两个指定节点的最近公共祖先。
//百度百科中最近公共祖先的定义为：
//对于有根树T的两个结点p、q，最近公共祖先表示为一个结点x，
//满足x是p、q的祖先且x的深度尽可能大（一个节点也可以是它自己的祖先）。
//所有节点的值都是唯一的。p、q 为不同节点且均存在于给定的二叉搜索树中。

//从根结点开始，第一个分叉点就是两个结点的最近公共祖先
//如果两个结点的值都小于根节点，那么两个结点都在根结点的左子树上
//把左子树的根节点设置为根结点，继续上述规程，反之搜索右子树

package main

import "fmt"

type TreeNode struct {
  Val   int
  Left  *TreeNode
  Right *TreeNode
}

func main() {
  p := &TreeNode{20, nil, nil}
  q := &TreeNode{60, nil, nil}
  phead :=
    &TreeNode{50,
      &TreeNode{30,
        p,
        &TreeNode{40, nil, nil}},
      &TreeNode{70,
        q,
        &TreeNode{80, nil, nil}},
    }

  fmt.Println(lowestCommonAncestor(phead, p, q))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  if p.Val < root.Val && q.Val < root.Val { //两个结点都小于根结点，继续向左搜索
    return lowestCommonAncestor(root.Left, p, q)
  }
  if p.Val > root.Val && q.Val > root.Val { //两个结点都大于根结点，继续向右搜索
    return lowestCommonAncestor(root.Right, p, q)
  }
  return root
}
