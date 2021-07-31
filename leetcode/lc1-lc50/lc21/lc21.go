package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
} //链表结点

func main() {
	ln13 := ListNode{Val: 4}
	ln12 := ListNode{Val: 2, Next: &ln13}
	ln11 := ListNode{Val: 1, Next: &ln12}

	ln23 := ListNode{Val: 4}
	ln22 := ListNode{Val: 3, Next: &ln23}
	ln21 := ListNode{Val: 1, Next: &ln22}

	tpLN := mergeTwoLists(&ln11, &ln21)

	for ; tpLN != nil; tpLN = tpLN.Next {
		fmt.Printf("%d,", tpLN.Val)
	}
}

//21-合并两个有序链表(21,88)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//使用两个指针标记两个链表中未处理的第一个结点
	//每次比较两个链表中未处理的第一个结点，把较小的结点连接到结果链表中

	var pLNTou2 *ListNode                 //头指针
	var pLNNow *ListNode                  //连接指针
	var tpLNl1, tpLNl2 *ListNode = l1, l2 //遍历指针

	if tpLNl1 == nil {
		return l2
	}
	if tpLNl2 == nil {
		return l1
	}

	//确定头结点
	if tpLNl1.Val >= tpLNl2.Val {
		pLNTou2 = tpLNl2
		pLNNow = tpLNl2
		tpLNl2 = tpLNl2.Next
	} else {
		pLNTou2 = tpLNl1
		pLNNow = tpLNl1
		tpLNl1 = tpLNl1.Next
	}

	for true {
		//其中一条链表已经到尾部
		if tpLNl1 == nil {
			pLNNow.Next = tpLNl2
			break
		}
		if tpLNl2 == nil {
			pLNNow.Next = tpLNl1
			break
		}
		//两条链表都没有到尾部
		if tpLNl1.Val >= tpLNl2.Val {
			pLNNow.Next = tpLNl2
			pLNNow = tpLNl2
			tpLNl2 = tpLNl2.Next
		} else {
			pLNNow.Next = tpLNl1
			pLNNow = tpLNl1
			tpLNl1 = tpLNl1.Next
		}
	}

	return pLNTou2
}
