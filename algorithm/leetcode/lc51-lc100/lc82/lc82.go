package main

import "fmt"

type ListNode struct {
  Val  int
  Next *ListNode
}

func main() {
  ln5 := ListNode{Val: 4}
  ln4 := ListNode{Val: 3, Next: &ln5}
  ln3 := ListNode{Val: 3, Next: &ln4}
  ln2 := ListNode{Val: 3, Next: &ln3}
  ln1 := ListNode{Val: 3, Next: &ln2}

  phead := deleteDuplicates(&ln1)
  for ; phead != nil; phead = phead.Next {
    fmt.Printf("%d,", phead.Val)
  }
}

//存在一个按升序排列的链表，给你这个链表的头节点head，
//删除链表中所有存在数字重复情况的节点，只保留原始链表中没有重复出现的数字。
//返回同样按升序排列的结果链表。
//链表中节点数目在范围[0,300]内;-100<=Node.val<=100;题目数据保证链表已经按升序排列

//82-删除排序链表中的重复元素II(82,83)
func deleteDuplicates(head *ListNode) *ListNode {
  var plink, pheadTemp *ListNode = nil, nil //遍历指针，临时头结点指针

  if head == nil {
    return head
  }

  pheadTemp = &ListNode{-200, head} //头结点可能被删除，安排一个临时头结点
  plink = pheadTemp                 //从临时头结点开始遍历

  for plink.Next != nil && plink.Next.Next != nil {
    if plink.Next.Val == plink.Next.Next.Val {
      //如果后面两个是一样的就跳过，直到找到下一个不一样的
      num := plink.Next.Val
      for plink.Next != nil && plink.Next.Val == num {
        plink.Next = plink.Next.Next
      }
    } else {
      plink = plink.Next
    }
  }

  return pheadTemp.Next
}
