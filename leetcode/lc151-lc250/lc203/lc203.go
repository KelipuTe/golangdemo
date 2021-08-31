package main

import "fmt"

type ListNode struct {
  Val  int
  Next *ListNode
}

func main() {
  ln6 := ListNode{6, nil}
  ln5 := ListNode{2, &ln6}
  ln4 := ListNode{4, &ln5}
  ln3 := ListNode{3, &ln4}
  ln2 := ListNode{2, &ln3}
  ln1 := ListNode{1, &ln2}

  // ln4 := ListNode{1, nil}
  // ln3 := ListNode{2, &ln4}
  // ln2 := ListNode{2, &ln3}
  // ln1 := ListNode{1, &ln2}

  phead := removeElements(&ln1, 2)

  for ; phead != nil; phead = phead.Next {
    fmt.Printf("%d,", phead.Val)
  }
}

//注意，开头连续，中间连续，末尾连续的情况

//203-移除链表元素
func removeElements(head *ListNode, val int) *ListNode {
  var phead, plink, pnow *ListNode = head, nil, nil

  for phead != nil {
    if phead.Val == val {
      phead = phead.Next
    } else {
      break
    }
  }

  if phead == nil {
    return nil
  }

  plink, pnow = phead, phead.Next
  for pnow != nil {
    if pnow.Val == val {
      pnow = pnow.Next
      continue
    }
    plink.Next = pnow
    plink = pnow
    pnow = pnow.Next
  }
  plink.Next = nil

  return phead
}
