package main

import "fmt"

//链表结点
type ListNode struct {
	Val  int
	Next *ListNode
}

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

//解：
//两个指针p1和p2，先让p1从头开始，往前遍历n个结点，然后p2从头开始同步遍历
//当p1遍历到链表尾部时，p2就指向倒数第n个结点处，需要错1位，让p2指向前一个结点才能进行删除

//删除链表的倒数第n个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	iListLen := 0 //链表长度
	pLNTou2 := head
	pLNWei3 := head
	tpLN := head //要删的结点的前1个结点

	if n < 1 {
		return head
	}

	//让pLNWei3指针领先tpLN指针n个结点
	for pLNWei3 != nil {
		pLNWei3 = pLNWei3.Next
		iListLen++
		if iListLen > n+1 {
			tpLN = tpLN.Next
		}
	}

	if iListLen > n {
		//n小于等于链表长度，删
		if tpLN.Next == pLNWei3 {
			//删的是尾节点
			tpLN.Next = nil
		} else {
			tpLN.Next = tpLN.Next.Next
		}
	} else if iListLen == n {
		//n等于链表长度，删第1个
		pLNTou2 = pLNTou2.Next
	}

	return pLNTou2
}
