package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
} //树结点

func main() {
	tn7 := TreeNode{7, nil, nil}
	tn6 := TreeNode{5, nil, nil}
	tn3 := TreeNode{5, &tn6, &tn7}
	tn2 := TreeNode{2, nil, nil}
	tn1 := TreeNode{2, &tn2, &tn3}

	fmt.Println(findSecondMinimumValue(&tn1))
}

//给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为2或0。
//如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个。
//找到所有节点中的第二小的值。如果第二小的值不存在的话，输出-1。

//树中节点数目在范围[1, 25]内；1<=Node.val<=(2^31)-1
//对于树中每个节点:root.val=min(root.left.val,root.right.val)

//671-二叉树中第二小的节点
func findSecondMinimumValue(root *TreeNode) int {
	//遍历一遍找到除了根结点最小的即可

	var iMinVal int = -1                        //设定一个临界值
	var psli1TN []*TreeNode = []*TreeNode{root} //队列

	if root == nil {
		return -1
	}

	for len(psli1TN) > 0 {
		tpTN := psli1TN[0]
		psli1TN = psli1TN[1:] //出队
		if tpTN.Val > root.Val {
			if iMinVal == -1 {
				iMinVal = tpTN.Val
			} else if tpTN.Val < iMinVal {
				iMinVal = tpTN.Val
			}
		}
		if tpTN.Left != nil {
			psli1TN = append(psli1TN, tpTN.Left)
		}
		if tpTN.Right != nil {
			psli1TN = append(psli1TN, tpTN.Right)
		}
	}

	return iMinVal
}
