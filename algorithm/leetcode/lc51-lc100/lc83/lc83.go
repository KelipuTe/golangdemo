package main

import "fmt"

type ListNode struct {
  Val  int
  Next *ListNode
}

func main() {
  ln5 := ListNode{Val: 3}
  ln4 := ListNode{Val: 3, Next: &ln5}
  ln3 := ListNode{Val: 2, Next: &ln4}
  ln2 := ListNode{Val: 1, Next: &ln3}
  ln1 := ListNode{Val: 1, Next: &ln2}

  phead := deleteDuplicates(&ln1)
  for ; phead != nil; phead = phead.Next {
    fmt.Printf("%d,", phead.Val)
  }
}

//存在一个按升序排列的链表，给你这个链表的头节点head，
//删除所有重复的元素，使每个元素只出现一次。
//返回同样按升序排列的结果链表。
//链表中节点数目在范围[0,300]内;-100<=Node.val<=100;题目数据保证链表已经按升序排列

//83-删除排序链表中的重复元素(82,83)
func deleteDuplicates(head *ListNode) *ListNode {
  var plink *ListNode = head

  if plink == nil {
    return nil
  }

  for plink.Next != nil {
    if plink.Val == plink.Next.Val {
      //如果后面一个和自己一样就跳过
      plink.Next = plink.Next.Next
    } else {
      plink = plink.Next
    }
  }

  return head
}
