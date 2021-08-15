package main

import "fmt"

func main() {
  ln2 := ListNode{2, nil}
  ln1 := ListNode{1, &ln2}
  ln2.Next = &ln1
  fmt.Println(hasCycle(&ln1))
}

type ListNode struct {
  Val  int
  Next *ListNode
}

//双指针
//使用两个指针，一快一慢，快指针一次走两个结点，慢指针一次走一个结点
//如果某个时刻两个指针指向的地址相等了，则存在环形结构

//141-环形链表
func hasCycle(head *ListNode) bool {
  if head == nil || head.Next == nil {
    return false
  }
  tpln1, tpln2 := head, head.Next

  for tpln1 != tpln2 {
    if tpln2 == nil || tpln2.Next == nil {
      return false
    }
    tpln1 = tpln1.Next
    tpln2 = tpln2.Next.Next
  }

  return true
}
