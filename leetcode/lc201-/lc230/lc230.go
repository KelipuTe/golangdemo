package main

import "fmt"

var isli1Num []int

//树结点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tn8 := TreeNode{1, nil, nil}
	tn5 := TreeNode{4, nil, nil}
	tn4 := TreeNode{2, &tn8, nil}
	tn3 := TreeNode{6, nil, nil}
	tn2 := TreeNode{3, &tn4, &tn5}
	tn1 := TreeNode{5, &tn2, &tn3}

	iRes := kthSmallest(&tn1, 3)
	fmt.Println(iRes)
}

//二叉搜索树中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	//二叉搜索树的中序遍历是升序序列
	if root == nil {
		return 0
	}
	isli1Num = []int{}
	fZhongXu(root)
	return isli1Num[k-1]
}

func fZhongXu(root *TreeNode) {
	if root.Left != nil {
		fZhongXu(root.Left)
	}
	isli1Num = append(isli1Num, root.Val)
	if root.Right != nil {
		fZhongXu(root.Right)
	}
}
