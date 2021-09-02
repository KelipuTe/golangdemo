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

  lnRes := removeNthFromEnd(&ln1, 3)
  for ; lnRes != nil; lnRes = lnRes.Next {
    fmt.Printf("%d,", lnRes.Val)
  }
}

//链表，双指针
//两个指针，先让第一个指针从头开始，往前遍历n个结点，然后第二个指针从头开始同步遍历
//当第一个指针遍历到链表尾部时，第二个指针就指向倒数第n个结点处

//19-删除链表的倒数第n个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
  var listLen int = 0                                   //链表长度
  var phead, ptail, pquery *ListNode = head, head, head //头指针，尾指针，遍历指针

  if n < 1 {
    return head
  }

  for ptail != nil {
    ptail = ptail.Next
    listLen++
    if listLen > n+1 { //需要错1位，让pquery指向前一个结点才能进行删除
      pquery = pquery.Next
    }
  }

  if listLen > n { //链表长度大于等于n，删倒数第n个
    if pquery.Next == ptail { //删的是尾节点
      pquery.Next = nil
    } else {
      pquery.Next = pquery.Next.Next
    }
  } else { //链表长度小于等于n，删第1个
    phead = phead.Next
  }

  return phead
}
