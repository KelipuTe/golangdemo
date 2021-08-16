package main

import "fmt"

func main() {
  ln4 := ListNode{4, nil}
  ln3 := ListNode{3, &ln4}
  ln2 := ListNode{2, &ln3}
  ln1 := ListNode{1, &ln2}
  ln4.Next = &ln2
  fmt.Println(detectCycle(&ln1))
}

type ListNode struct {
  Val  int
  Next *ListNode
}

//给定一个链表，返回链表开始入环的第一个节点。如果链表无环，则返回null。
//为了表示给定链表中的环，我们使用整数pos来表示链表尾连接到链表中的位置（索引从0开始）。
//如果pos是-1，则在该链表中没有环。注意：pos不作为参数进行传递，仅仅是为了标识链表的实际情况。
//链表中节点的数目范围是[0,10^4];-10^5<=Node.val<=10^5;pos为-1或者链表中的一个有效索引。
//说明：不允许修改给定的链表。进阶：用O(1)空间解决此题。

//双指针变种
//使用一个步长为1的慢指针和一个步长为2的快指针，从头部开始同步遍历。
//如果某个时刻，两个指针指向了相同的地址，则存在环形结构。
//如果存在环形结构，快指针会绕1圈或者几圈之后，从后面追上慢指针。
//这样可以判断有无环形结构，但是无法确认入环点的位置。

//假设入环点（不包括）之前的结点个数为a，
//入环点（包括）到相遇点（包括）之间的结点个数为b，
//相遇点（不包括）再次到入环点（不包括）之间的结点个数为c。
//根据快慢指针走过的结点个数，可以得到：`2*(a+b)=a+n*(b+c)+b`。推导得出：`a=(n-1)*b+n*c=(n-1)*(b+c)+c`。
//根据推导出的公式，我们可以得出，如果此时，再使用一个步长为1的慢指针，从头开始同步遍历。
//那么这个慢指针会和之前的慢指针在入环点相遇。

//142-环形链表II(141,142)
func detectCycle(head *ListNode) *ListNode {
  if head == nil {
    return nil
  }

  tpln1, tpln2 := head, head        //慢指针，快指针
  tpln3, btpln3Start := head, false //第二个慢指针

  for true {
    if tpln2 == nil || tpln2.Next == nil {
      return nil //快指针能走到头，说明没有环
    }
    tpln1 = tpln1.Next
    tpln2 = tpln2.Next.Next
    if btpln3Start {
      tpln3 = tpln3.Next
    }
    if tpln1 == tpln2 {
      btpln3Start = true //第二个慢指针开始遍历
    }
    if btpln3Start && tpln1 == tpln3 {
      return tpln1
    }
  }

  return nil
}
