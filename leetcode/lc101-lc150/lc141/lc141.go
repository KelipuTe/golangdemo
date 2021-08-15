package main

import "fmt"

func main() {
  ln1 := ListNode{2, nil}
  fmt.Println(hasCycle(&ln1))
}

type ListNode struct {
  Val  int
  Next *ListNode
}

//双指针
//使用两个指针，一快一慢，快指针一次走两个结点，慢指针一次走一个结点
//如果某个时刻两个指针指向的地址相等了，存在环形结构

//141-环形链表
func hasCycle(head *ListNode) bool {
if head==nil{
  return false
}
if head.Next==nil{
  return false
}
tpln1=head.Next

  for true {
    tpln1 := head.Next.Next
    tpln2 := head.Next
    if
  }
}
