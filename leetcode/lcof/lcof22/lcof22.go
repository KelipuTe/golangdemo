package main

import "fmt"

type ListNode struct {
  Val  int
  Next *ListNode
}

func main() {
  phead := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
  pres := getKthFromEnd(phead, 3)
  for ; pres != nil; pres = pres.Next {
    fmt.Printf("%d,", pres.Val)
  }
}

//链表，双指针
//两个指针，先让第一个指针从头开始，往前遍历n个结点，然后第二个指针从头开始同步遍历
//当第一个指针遍历到链表尾部时，第二个指针就指向倒数第n个结点处

//剑指Offer22-链表中倒数第k个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
  var listLen int = 0                      //链表长度
  var ptail, pquery *ListNode = head, head //尾指针，遍历指针

  if k < 1 {
    return head
  }

  for ptail != nil {
    ptail = ptail.Next
    listLen++
    if listLen > k {
      pquery = pquery.Next
    }
  }

  if listLen > k { //链表长度大于等于n，返回第n个
    return pquery
  } else { //链表长度小于等于n，返回第1个
    return head
  }
}
