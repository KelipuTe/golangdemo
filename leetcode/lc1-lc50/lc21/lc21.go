package main

import "fmt"

type ListNode struct {
  Val  int
  Next *ListNode
}

func main() {
  ln13 := ListNode{Val: 4}
  ln12 := ListNode{Val: 2, Next: &ln13}
  ln11 := ListNode{Val: 1, Next: &ln12}

  ln23 := ListNode{Val: 4}
  ln22 := ListNode{Val: 3, Next: &ln23}
  ln21 := ListNode{Val: 1, Next: &ln22}

  phead := mergeTwoLists(&ln11, &ln21)

  for ; phead != nil; phead = phead.Next {
    fmt.Printf("%d,", phead.Val)
  }
}

//链表，双指针
//使用两个指针标记两个链表中未处理的第一个结点
//每次比较两个链表中未处理的第一个结点，把较小的结点连接到结果链表中

//21-合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
  var phead, plink *ListNode      //头指针，连接指针
  var pl1, pl2 *ListNode = l1, l2 //遍历指针

  if pl1 == nil {
    return l2
  }
  if pl2 == nil {
    return l1
  }

  //确定头结点
  if pl1.Val >= pl2.Val {
    phead = pl2
    plink = pl2
    pl2 = pl2.Next
  } else {
    phead = pl1
    plink = pl1
    pl1 = pl1.Next
  }

  for true {
    //其中一条链表已经到尾部
    if pl1 == nil {
      plink.Next = pl2
      break
    }
    if pl2 == nil {
      plink.Next = pl1
      break
    }
    //两条链表都没有到尾部
    if pl1.Val >= pl2.Val {
      plink.Next = pl2
      plink = pl2
      pl2 = pl2.Next
    } else {
      plink.Next = pl1
      plink = pl1
      pl1 = pl1.Next
    }
  }

  return phead
}
