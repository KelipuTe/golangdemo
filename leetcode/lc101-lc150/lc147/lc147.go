package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
} //链表结点

func main() {
	ln15 := ListNode{Val: 4}
	ln14 := ListNode{Val: 3, Next: &ln15}
	ln13 := ListNode{Val: 6, Next: &ln14}
	ln12 := ListNode{Val: 2, Next: &ln13}
	ln11 := ListNode{Val: 1, Next: &ln12}

	tpLN := insertionSortList(&ln11)
	for ; tpLN != nil; tpLN = tpLN.Next {
		fmt.Printf("%d,", tpLN.Val)
	}
}

//148-对链表进行插入排序
func insertionSortList(head *ListNode) *ListNode {
	var tpLNHead *ListNode = head     //默认头结点有序
	var pLNHead *ListNode = head.Next //原链表剩余部分

	tpLNHead.Next = nil //断开头结点后面的链表

	//处理剩下的部分
	for pLNHead != nil {
		//把结点取出来，断开结点后面的链表
		tpLN := pLNHead
		pLNHead = pLNHead.Next
		tpLN.Next = nil
		if tpLN.Val < tpLNHead.Val {
			//比头结点小
			tpLN.Next = tpLNHead
			tpLNHead = tpLN
		} else {
			//遍历有序链表
			tpLN2 := tpLNHead
			for tpLN2 != nil {
				if tpLN2.Next == nil {
					//链表尾部
					tpLN2.Next = tpLN
					break
				}
				if tpLN2.Val <= tpLN.Val && tpLN2.Next.Val > tpLN.Val {
					//链表中间
					tpLN.Next = tpLN2.Next
					tpLN2.Next = tpLN
					break
				}
				tpLN2 = tpLN2.Next
			}
		}
	}

	return tpLNHead
}
