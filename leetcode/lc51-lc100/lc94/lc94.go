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

	fmt.Println(inorderTraversal(&tn1))
}

//树中节点数目在范围[0,100]内；-100<=Node.val<=100

var isli1QueryRes []int //遍历结果

//94、二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	isli1QueryRes = []int{}
	if root == nil {
		return isli1QueryRes
	}
	fZhongXuQuery(root)
	return isli1QueryRes
}

func fZhongXuQuery(root *TreeNode) {
	if root.Left != nil {
		fZhongXuQuery(root.Left)
	}
	isli1QueryRes = append(isli1QueryRes, root.Val)
	if root.Right != nil {
		fZhongXuQuery(root.Right)
	}
}
