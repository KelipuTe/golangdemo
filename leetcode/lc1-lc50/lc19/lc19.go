package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
} //链表结点

func main() {
	ln5 := ListNode{Val: 5}
	ln4 := ListNode{Val: 4, Next: &ln5}
	ln3 := ListNode{Val: 3, Next: &ln4}
	ln2 := ListNode{Val: 2, Next: &ln3}
	ln1 := ListNode{Val: 1, Next: &ln2}

	for tp := removeNthFromEnd(&ln1, 3); tp != nil; tp = tp.Next {
		fmt.Printf("%d,", tp.Val)
	}
}

//删除链表的倒数第n个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var iListLen int = 0           //链表长度
	var pLNTou2 *ListNode = head   //头指针
	var pLNWei3 *ListNode = head   //尾指针
	var tpLNQuery *ListNode = head //遍历指针

	if n < 1 {
		return head
	}

	//两个指针，先让pLNWei3从头开始，往前遍历n个结点，然后tpLNQuery从头开始同步遍历
	//当pLNWei3遍历到链表尾部时，tpLNQuery就指向倒数第n个结点处
	for pLNWei3 != nil {
		pLNWei3 = pLNWei3.Next
		iListLen++
		if iListLen > n+1 {
			//需要错1位，让tpLNQuery指向前一个结点才能进行删除
			tpLNQuery = tpLNQuery.Next
		}
	}

	if iListLen > n {
		//n小于等于链表长度，删
		if tpLNQuery.Next == pLNWei3 {
			//删的是尾节点
			tpLNQuery.Next = nil
		} else {
			tpLNQuery.Next = tpLNQuery.Next.Next
		}
	} else if iListLen == n {
		//n等于链表长度，删第1个
		pLNTou2 = pLNTou2.Next
	}

	return pLNTou2
}
