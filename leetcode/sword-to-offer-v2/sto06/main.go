package main

import "fmt"

// 剑指 Offer 06. 从尾到头打印链表
// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

func main() {
  // p1head := &ListNode{1, &ListNode{3, &ListNode{2, nil}}}
  p1head := &ListNode{1, &ListNode{3, &ListNode{2, &ListNode{4, &ListNode{5, &ListNode{8, nil}}}}}}
  fmt.Println(reversePrint(p1head))
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
  Val  int
  Next *ListNode
}

func reversePrint(head *ListNode) []int {
  if nil == head {
    return []int{}
  }

  t0sli1num := make([]int, 0, 16)
  j := 0

  p1now := head
  for nil != p1now {
    t0sli1num = append(t0sli1num, p1now.Val)
    j++
    if nil != p1now.Next {
      p1now = p1now.Next
    } else {
      break
    }
  }

  for i := 0; i < (j >> 1); i++ {
    t0sli1num[i], t0sli1num[j-1-i] = t0sli1num[j-1-i], t0sli1num[i]
  }

  return t0sli1num
}
