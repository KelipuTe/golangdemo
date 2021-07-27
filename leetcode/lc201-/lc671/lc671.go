package main

import "fmt"

//树结点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tn7 := TreeNode{7, nil, nil}
	tn6 := TreeNode{5, nil, nil}
	tn3 := TreeNode{5, &tn6, &tn7}
	tn2 := TreeNode{2, nil, nil}
	tn1 := TreeNode{2, &tn2, &tn3}

	iRes := findSecondMinimumValue(&tn1)
	fmt.Println(iRes)
}

//二叉树中第二小的节点
//给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为2或0。
//如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个。
//即root.val = min(root.left.val, root.right.val)总成立。
//找到所有节点中的第二小的值。如果第二小的值不存在的话，输出-1。

func findSecondMinimumValue(root *TreeNode) int {
	//遍历一遍找到除了根节点最小的即可
	if root == nil {
		return -1
	}

	var psli1TN []*TreeNode = []*TreeNode{root}
	var iRes int = -1

	for len(psli1TN) > 0 {
		tpTN := psli1TN[0]
		psli1TN = psli1TN[1:]
		if tpTN.Val > root.Val {
			if iRes == -1 {
				iRes = tpTN.Val
			} else if tpTN.Val < iRes {
				iRes = tpTN.Val
			}
		}
		if tpTN.Left != nil {
			psli1TN = append(psli1TN, tpTN.Left)
		}
		if tpTN.Right != nil {
			psli1TN = append(psli1TN, tpTN.Right)
		}
	}

	return iRes
}
