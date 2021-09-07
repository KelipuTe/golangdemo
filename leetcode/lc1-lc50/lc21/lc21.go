//lc21-合并两个有序链表
//[链表][双指针]

//将两个升序链表合并为一个新的升序链表并返回。
//新链表是通过拼接给定的两个链表的所有节点组成的。

//使用两个指针标记两个链表中未处理的第一个结点
//每次比较两个链表中未处理的第一个结点，把较小的结点连接到结果链表中

package main

import "fmt"

type ListNode struct {
  Val  int
  Next *ListNode
}

func main() {
  p1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
  p2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
  phead := mergeTwoLists(p1, p2)

  for ; phead != nil; phead = phead.Next {
    fmt.Printf("%d,", phead.Val)
  }
}

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
    phead, plink = pl2, pl2
    pl2 = pl2.Next
  } else {
    phead, plink = pl1, pl1
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
