package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
} //树结点

func main() {
	tn8 := TreeNode{1, nil, nil}
	tn5 := TreeNode{4, nil, nil}
	tn4 := TreeNode{2, &tn8, nil}
	tn3 := TreeNode{6, nil, nil}
	tn2 := TreeNode{3, &tn4, &tn5}
	tn1 := TreeNode{5, &tn2, &tn3}

	fmt.Println(postorderTraversal(&tn1))
}

//树中节点数目在范围[0,100]内
//-100<=Node.val<=100

var isli1QueryRes []int //遍历结果

//145-二叉树的后序遍历(94,144,145)
func postorderTraversal(root *TreeNode) []int {
	isli1QueryRes = []int{}
	if root == nil {
		return isli1QueryRes
	}
	fHouXuQuery(root)
	return isli1QueryRes
}

func fHouXuQuery(root *TreeNode) {
	if root.Left != nil {
		fHouXuQuery(root.Left)
	}
	if root.Right != nil {
		fHouXuQuery(root.Right)
	}
	isli1QueryRes = append(isli1QueryRes, root.Val)
}
