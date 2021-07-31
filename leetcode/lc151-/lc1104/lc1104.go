package main

import "fmt"

func main() {
	fmt.Println(pathInZigZagTree(26))
}

//在一棵无限的二叉树上，每个节点都有两个子节点。
//数值从1开始，在奇数行中，按从左到右的顺序进行标记；而偶数行中，按从右到左的顺序进行标记。
//给树上某一个节点的标号label，返回从根节点到该标号的路径，该路径是由途经的节点标号所组成的。
//1<=label<=10^6

//1104-二叉树寻路
func pathInZigZagTree(label int) []int {
	//完全二叉树，下标为n的子结点的父结点的下标是n/2取整
	//奇数行n，节点从左到右标记，label在下标label的位置
	//偶数行n，节点从右到左标记，label在下标2^(n-1)+2^n-1-label的位置

	iRow := 1                     //行数
	iRowFirstNum := 1             //行开始的第一个数值
	var iRealIndex int = label    //label的真实位置
	var isli1Path []int = []int{} //结果路径

	for iRowFirstNum*2 <= label {
		iRow++
		iRowFirstNum *= 2
	}
	if iRow%2 == 0 {
		iRealIndex = 1<<(iRow-1) + 1<<(iRow) - 1 - label
	}

	for iRow > 0 {
		//判断结点的真实位置
		if iRow%2 == 0 {
			isli1Path = append(isli1Path, 1<<(iRow-1)+1<<(iRow)-1-iRealIndex)
		} else {
			isli1Path = append(isli1Path, iRealIndex)
		}
		//父结点的位置
		iRow--
		iRealIndex = iRealIndex >> 1
	}
	//遍历结果反续
	for ii, iLen := 0, len(isli1Path); ii < iLen/2; ii++ {
		isli1Path[ii], isli1Path[iLen-1-ii] = isli1Path[iLen-1-ii], isli1Path[ii]
	}

	return isli1Path
}
