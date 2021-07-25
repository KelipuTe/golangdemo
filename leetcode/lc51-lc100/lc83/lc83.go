package main

import "fmt"

//链表结点
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ln5 := ListNode{Val: 3}
	ln4 := ListNode{Val: 3, Next: &ln5}
	ln3 := ListNode{Val: 2, Next: &ln4}
	ln2 := ListNode{Val: 1, Next: &ln3}
	ln1 := ListNode{Val: 1, Next: &ln2}

	tpln := deleteDuplicates(&ln1)
	for tpln != nil {
		fmt.Printf("%d,", tpln.Val)
		tpln = tpln.Next
	}
}

//删除排序链表中的重复元素，重复元素只保留一个结点
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		//链表为空或者只有一个结点
		return head
	}

	var tplnLink *ListNode = head       //连接指针
	var tplnCheck *ListNode = head      //校验指针
	var tplnQuery *ListNode = head.Next //遍历指针

	//遍历到最后一个结点
	for tplnQuery.Next != nil {
		if tplnQuery.Val != tplnCheck.Val {
			//校验结点不重复，连接
			tplnLink.Next = tplnQuery
			tplnLink = tplnQuery
			tplnCheck = tplnQuery
		}
		tplnQuery = tplnQuery.Next
	}

	//处理最后一个结点
	if tplnQuery.Val != tplnCheck.Val {
		tplnLink.Next = tplnQuery
	} else {
		//断开连接结点后面的链表
		tplnLink.Next = nil
	}
	return head
}
