package main

import "fmt"

func main() {
  ln5 := ListNode{5, nil}
  ln4 := ListNode{4, &ln5}
  ln3 := ListNode{3, &ln4}
  ln2 := ListNode{2, &ln3}
  ln1 := ListNode{1, &ln2}

  tpln := reverseList(&ln1)

  for tpln != nil {
    fmt.Println(tpln.Val)
    tpln = tpln.Next
  }
}

type ListNode struct {
  Val  int
  Next *ListNode
}

//给你单链表的头节点head，请反转链表，并返回反转后的链表。
//链表中节点的数目范围是[0,5000];-5000<=Node.val<=5000

//206-反转链表
func reverseList(head *ListNode) *ListNode {
  var plnP, plnT, plnN, plnHeadNew *ListNode

  if head == nil || head.Next == nil {
    return head //空链表或者只有一个结点
  }

  plnP, plnT, plnN = head, head.Next, head.Next.Next //前一个结点，当前结点，下一个结点
  for plnT != nil {
    if plnT.Next == nil {
      plnHeadNew = plnT //尾结点的next为空，反转之前就要判断，并设置为新得头结点
    }
    plnT.Next = plnP //反转
    plnP = plnT
    plnT = plnN
    if plnN != nil {
      plnN = plnN.Next
    }
  }
  head.Next = nil //处理原头结点

  return plnHeadNew
}
