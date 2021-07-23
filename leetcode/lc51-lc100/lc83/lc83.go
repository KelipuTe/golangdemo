package main

import "fmt"

//链表结点
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ln5 := ListNode{Val: 3}
	// ln4 := ListNode{Val: 3, Next: &ln5}
	// ln3 := ListNode{Val: 2, Next: &ln4}
	// ln2 := ListNode{Val: 1, Next: &ln3}
	// ln1 := ListNode{Val: 1, Next: &ln2}

	tpln := deleteDuplicates(&ln5)
	for tpln != nil {
		fmt.Printf("%d,", tpln.Val)
		tpln = tpln.Next
	}
}

//删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tplnLink := head
	iCheckNum := head.Val
	tplnQuery := head.Next
	//处理到最后一个
	for tplnQuery.Next != nil {
		if tplnQuery.Val != iCheckNum {
			tplnLink.Next = tplnQuery
			tplnLink = tplnLink.Next
			iCheckNum = tplnQuery.Val
		}
		tplnQuery = tplnQuery.Next
	}
	//处理最后一个
	if tplnQuery.Val != iCheckNum {
		tplnLink.Next = tplnQuery
	} else {
		tplnLink.Next = nil
	}
	return head
}
