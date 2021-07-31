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

	fmt.Println(kthSmallest(&tn1, 3))
}

//查找二叉搜索树中第k个最小元素（从1开始计数）
//树中的节点数为n;1<=k<=n<=10^4;0<=Node.val<=10^4

var isli1QueryRes []int //遍历结果

//230-二叉搜索树中第K小的元素(94,98,144,145)
func kthSmallest(root *TreeNode, k int) int {
	//二叉搜索树的中序遍历是升序序列，遍历二叉搜索树然后取第k个值

	isli1QueryRes = []int{}
	fZhongXuQuery(root)
	return isli1QueryRes[k-1]
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
