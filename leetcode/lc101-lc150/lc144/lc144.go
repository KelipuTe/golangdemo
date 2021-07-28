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

	fmt.Println(preorderTraversal(&tn1))
}

//树中节点数目在范围[0,100]内；-100<=Node.val<=100

var isli1QueryRes []int //遍历结果

//144、二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	isli1QueryRes = []int{}
	if root == nil {
		return isli1QueryRes
	}
	fQianXuQuery(root)
	return isli1QueryRes
}

func fQianXuQuery(root *TreeNode) {
	isli1QueryRes = append(isli1QueryRes, root.Val)
	if root.Left != nil {
		fQianXuQuery(root.Left)
	}
	if root.Right != nil {
		fQianXuQuery(root.Right)
	}
}
