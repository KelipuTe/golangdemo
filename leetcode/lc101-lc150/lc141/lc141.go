package main

import "fmt"

func main() {
  ln4 := ListNode{4, nil}
  ln3 := ListNode{3, &ln4}
  ln2 := ListNode{2, &ln3}
  ln1 := ListNode{1, &ln2}
  ln4.Next = &ln2
  fmt.Println(hasCycle(&ln1))
}

type ListNode struct {
  Val  int
  Next *ListNode
}

//给定一个链表，判断链表中是否有环。如果链表中存在环，则返回true。否则，返回false。
//如果链表中有某个节点，可以通过连续跟踪next指针再次到达，则链表中存在环。
//为了表示给定链表中的环，我们使用整数pos来表示链表尾连接到链表中的位置（索引从0开始）。
//如果pos是-1，则在该链表中没有环。注意：pos不作为参数进行传递，仅仅是为了标识链表的实际情况。
//链表中节点的数目范围是[0,10^4];-10^5<=Node.val<=10^5;pos为-1或者链表中的一个有效索引。
//进阶：用O(1)内存解决此问题

//双指针
//使用一个步长为1的慢指针和一个步长为2的快指针，从头部开始同时向后遍历。
//如果某个时刻，两个指针指向了相同的地址，则存在环形结构。
//如果存在环形结构，快指针会绕1圈或者几圈之后，从后面追上慢指针。

//141-环形链表(141,142)
func hasCycle(head *ListNode) bool {
  if head == nil {
    return false
  }

  tpln1, tpln2 := head, head //慢指针，快指针

  for true {
    if tpln2 == nil || tpln2.Next == nil {
      return false //快指针能走到头，说明没有环
    }
    tpln1 = tpln1.Next
    tpln2 = tpln2.Next.Next
    if tpln1 == tpln2 {
      return true
    }
  }

  return false
}
