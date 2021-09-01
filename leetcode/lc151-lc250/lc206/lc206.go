package main

import "fmt"

type ListNode struct {
  Val  int
  Next *ListNode
}

func main() {
  ln5 := ListNode{5, nil}
  ln4 := ListNode{4, &ln5}
  ln3 := ListNode{3, &ln4}
  ln2 := ListNode{2, &ln3}
  ln1 := ListNode{1, &ln2}

  phead := reverseList(&ln1)
  for ; phead != nil; phead = phead.Next {
    fmt.Printf("%d,", phead.Val)
  }
}

//给你单链表的头节点head，请反转链表，并返回反转后的链表。
//链表中节点的数目范围是[0,5000];-5000<=Node.val<=5000

//206-反转链表
func reverseList(head *ListNode) *ListNode {
  var pprevious, pnow, pnext *ListNode = nil, nil, nil //前一个结点，当前结点，下一个结点

  pnow = head
  for pnow != nil {
    pnext = pnow.Next     //原来的下一个结点
    pnow.Next = pprevious //反转，第一次的时候是nil
    pprevious = pnow      //当前结点变成前一个结点
    pnow = pnext          //原来的下一个结点变成当前结点
  }

  return pprevious
}
