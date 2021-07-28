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
	tn4 := TreeNode{10, &tn8, nil}
	tn3 := TreeNode{6, nil, nil}
	tn2 := TreeNode{3, &tn4, &tn5}
	tn1 := TreeNode{5, &tn2, &tn3}

	fmt.Println(isValidBST(&tn1))
}

//给定一个二叉树，判断其是否是一个有效的二叉搜索树

//二叉搜索树的中序遍历是升序序列，遍历二叉搜索树然后判断是不是升序序列

var isli1QueryRes []int //遍历结果

//98、验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	isli1QueryRes = []int{}
	fZhongXuQuery(root)
	for iIndex := 1; iIndex < len(isli1QueryRes); iIndex++ {
		if isli1QueryRes[iIndex] <= isli1QueryRes[iIndex-1] {
			return false
		}
	}
	return true
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
