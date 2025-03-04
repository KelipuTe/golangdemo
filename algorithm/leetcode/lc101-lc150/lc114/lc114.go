package main

func main() {
  tn7 := TreeNode{3, nil, nil}
  tn6 := TreeNode{4, nil, nil}
  tn5 := TreeNode{4, nil, nil}
  tn4 := TreeNode{3, nil, nil}
  tn3 := TreeNode{2, &tn6, &tn7}
  tn2 := TreeNode{2, &tn4, &tn5}
  tn1 := TreeNode{1, &tn2, &tn3}

  flatten(&tn1)
}

type TreeNode struct {
  Val   int
  Left  *TreeNode
  Right *TreeNode
} // 树结点

//给你二叉树的根结点root ，请你将它展开为一个单链表
//展开后的单链表应该同样使用TreeNode ，其中right子指针指向链表中下一个结点，而左子指针始终为null
//展开后的单链表应该与二叉树先序遍历顺序相同

//使用一个指针数组保存前序遍历结果，指针指向二叉树的结点
//然后按照数组顺序依次修改每一个结点的左右指针

var ptnsli1Res []*TreeNode //前序遍历结果

//114-二叉树展开为链表
func flatten(root *TreeNode) {
  ptnsli1Res = []*TreeNode{}
  xian1xu4bian4li4(root)
  for ii := 1; ii < len(ptnsli1Res); ii++ {
    tptn1, tptn2 := ptnsli1Res[ii-1], ptnsli1Res[ii]
    tptn1.Left, tptn1.Right = nil, tptn2
  }
}

func xian1xu4bian4li4(root *TreeNode) {
  if root == nil {
    return
  }
  ptnsli1Res = append(ptnsli1Res, root)
  xian1xu4bian4li4(root.Left)
  xian1xu4bian4li4(root.Right)
}
