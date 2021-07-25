package main

import "fmt"

//链表结点
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ln5 := ListNode{Val: 4}
	ln4 := ListNode{Val: 3, Next: &ln5}
	ln3 := ListNode{Val: 3, Next: &ln4}
	ln2 := ListNode{Val: 3, Next: &ln3}
	ln1 := ListNode{Val: 3, Next: &ln2}

	tpln := deleteDuplicates(&ln1)
	for tpln != nil {
		fmt.Printf("%d,", tpln.Val)
		tpln = tpln.Next
	}
}

//删除排序链表中的重复元素，删除重复元素的所有结点
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		//链表为空或者只有一个结点
		return head
	}

	var tplnHead *ListNode = nil        //头指针
	var tplnLink *ListNode = nil        //连接指针
	var tplnCheck *ListNode = head      //校验指针
	var iNumCount int = 1               //校验计数
	var tplnQuery *ListNode = head.Next //遍历指针

	for tplnQuery != nil {
		if tplnQuery.Val == tplnCheck.Val {
			//校验结点重复
			iNumCount++
		} else {
			if iNumCount == 1 {
				//校验结点不重复，连接
				if tplnHead == nil {
					tplnHead = tplnCheck
					tplnLink = tplnCheck
					tplnCheck = tplnQuery
					iNumCount = 1
				} else {
					tplnLink.Next = tplnCheck
					tplnLink = tplnCheck
					tplnCheck = tplnQuery
					iNumCount = 1
				}
			} else {
				//校验结点重复，跳过
				tplnCheck = tplnQuery
				iNumCount = 1
			}
		}
		tplnQuery = tplnQuery.Next
	}

	//判断最后一个校验结点
	if iNumCount == 1 {
		//校验结点不重复，连接
		if tplnHead == nil {
			tplnHead = tplnCheck
		} else {
			tplnLink.Next = tplnCheck
		}
	} else {
		//校验结点重复，断开连接结点后面的链表
		if tplnHead != nil {
			tplnLink.Next = nil
		}
	}
	return tplnHead
}
