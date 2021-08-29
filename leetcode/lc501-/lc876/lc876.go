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
  lnRes := middleNode(&ln1)
  fmt.Println(lnRes.Val)
}

//给定一个头结点为head的非空单链表，返回链表的中间结点。
//如果有两个中间结点，则返回第二个中间结点。
//给定链表的结点数介于1和100之间。

//链表，双指针
//用一快一慢两个指针遍历链表，快指针一个一次走两个位置，慢指针一个一次走一个位置
//当快指针走到结尾时，慢指针就在中间位置

//876-链表的中间结点
func middleNode(head *ListNode) *ListNode {
  var pslow, pfast *ListNode = head, head

  for pfast.Next != nil {
    pslow = pslow.Next
    pfast = pfast.Next
    if pfast.Next != nil {
      pfast = pfast.Next
    }
  }

  return pslow
}
