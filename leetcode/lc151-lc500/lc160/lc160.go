package main

import (
  "fmt"
)

func main() {
  ln22 := ListNode{22, nil}
  ln21 := ListNode{21, &ln22}
  ln12 := ListNode{12, nil}
  ln11 := ListNode{11, &ln12}
  fmt.Println(getIntersectionNode(&ln11, &ln21))

  ln32 := ListNode{32, nil}
  ln41 := ListNode{41, &ln32}
  ln31 := ListNode{31, &ln32}
  fmt.Println(getIntersectionNode(&ln31, &ln41))
}

type ListNode struct {
  Val  int
  Next *ListNode
}

//给你两个单链表的头节点headA和headB，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回null。
//题目数据保证整个链式结构中不存在环。注意，函数返回结果后，链表必须保持其原始结构。
//listA中节点数目为m;listB中节点数目为n;0<=m,n<=3*10^4;1<=Node.val<=10^5;0<=skipA<=m;0<=skipB<=n;
//如果listA和listB没有交点，intersectVal为0。如果listA和listB有交点，intersectVal==listA[skipA+1]==listB[skipB+1]

//双指针，数学
//两个指针pa,pb分别从headA和headB开始同步向后遍历，
//当指针pa遍历到headA的末尾时，就再从headB开始向后遍历，指针pb同理
//当两个指针相遇时，就是两个链表的交点，当两个指针同时为空时，证明两个链表不存在交点
//相交的情况，pa和pb都走了m+n个结点，也就是说从交点x开始直到末尾，pa和pb都在一起
//不相交的情况，pa和pb在都走了m+n个结点后，分别到达对方的结尾

//160-相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
  if headA == nil || headB == nil {
    return nil
  }

  tplnA, tplnB := headA, headB
  for true {
    if tplnA == tplnB {
      return tplnA
    }
    if tplnA == nil && tplnB == nil {
      return nil
    }
    if tplnA != nil {
      tplnA = tplnA.Next
    } else {
      tplnA = headB
    }
    if tplnB != nil {
      tplnB = tplnB.Next
    } else {
      tplnB = headA
    }
  }

  return nil
}
