//lc19-删除链表的倒数第n个结点
//[单链表][双指针]

//给一个链表，删除链表的倒数第n个结点，并且返回链表的头结点。
//进阶：尝试使用一趟扫描实现。

//详见/shu4ju4jie2gou4/lian4biao3/dan1lian4biao3.DelNthFromTail()

package main

import "fmt"

type ListNode struct {
  Val  int
  Next *ListNode
}

func main() {
  phead := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
  phead = removeNthFromEnd(phead, 3)
  for pquery := phead; pquery != nil; pquery = pquery.Next {
    fmt.Printf("%d,", pquery.Val)
  }
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
  var pheadTemp *ListNode = &ListNode{0, head}           //临时头结点
  var pquery, pprevious *ListNode = pheadTemp, pheadTemp //遍历指针，倒数第n+1个结点

  if pquery.Next == nil {
    return pheadTemp.Next
  }

  for i := 1; i <= n+1; i++ { //要删除倒数第n个结点，需要倒数第n+1个结点
    pquery = pquery.Next
    if pquery == nil { //如果链表长度不够倒数，则删除第一个结点
      pprevious.Next = pprevious.Next.Next

      return pheadTemp.Next
    }
  }
  for pquery != nil { //同步遍历
    pquery = pquery.Next
    pprevious = pprevious.Next
  }

  pprevious.Next = pprevious.Next.Next //删除倒数第n个结点

  return pheadTemp.Next
}
