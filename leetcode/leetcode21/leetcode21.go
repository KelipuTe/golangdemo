package main

import "fmt"

//链表结点
type ListNode struct {
	Val  int
	Next *ListNode
}

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

//合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var pLNTou2, pLNNow *ListNode
	tpl1 := l1
	tpl2 := l2
	if tpl1 == nil {
		return l2
	}
	if tpl2 == nil {
		return l1
	}
	//确定头
	if tpl1.Val >= tpl2.Val {
		pLNTou2 = tpl2
		pLNNow = tpl2
		tpl2 = tpl2.Next
	} else {
		pLNTou2 = tpl1
		pLNNow = tpl1
		tpl1 = tpl1.Next
	}

	for true {
		//其中一条链表已经到尾部
		if tpl1 == nil {
			pLNNow.Next = tpl2
			break
		}
		if tpl2 == nil {
			pLNNow.Next = tpl1
			break
		}
		//两条链表都没有到尾部
		if tpl1.Val >= tpl2.Val {
			pLNNow.Next = tpl2
			pLNNow = tpl2
			tpl2 = tpl2.Next
		} else {
			pLNNow.Next = tpl1
			pLNNow = tpl1
			tpl1 = tpl1.Next
		}
	}

	return pLNTou2
}
