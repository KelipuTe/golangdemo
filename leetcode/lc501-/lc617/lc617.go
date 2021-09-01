package main

type TreeNode struct {
  Val   int
  Left  *TreeNode
  Right *TreeNode
}

func main() {

}

//给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
//你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，
//那么将他们的值相加作为节点合并后的新值，否则不为NULL的节点将直接作为新二叉树的节点。
//注意:合并必须从两个树的根节点开始。

//二叉树，递归

//617-合并二叉树
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
  //把root2和到root1上

  //两棵树，有一棵这个结点是空的
  if root1 == nil {
    return root2
  }
  if root2 == nil {
    return root1
  }
  //两棵树，这个结点没有空的
  root1.Val += root2.Val
  //递归处理这个结点的左右子结点
  root1.Left = mergeTrees(root1.Left, root2.Left)
  root1.Right = mergeTrees(root1.Right, root2.Right)

  return root1
}
