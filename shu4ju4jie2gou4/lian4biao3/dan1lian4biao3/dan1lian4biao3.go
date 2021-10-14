//单链表
package main

import "fmt"

type ListNode struct {
  Value int
  Next  *ListNode
}

func main() {
  var phead *ListNode = ListInit() //头结点

  AddNodeToTail(phead, 1)
  AddNodeListToTail(phead, 2, 3, 4, 5, 6, 7, 8)

  PrintList(phead)

  DelNthFromTail(phead, 8)
  DelNthFromTail(phead, 1)

  PrintList(phead)
}

//初始化
func ListInit() *ListNode {
  return &ListNode{0, nil}
}

//输出链表
func PrintList(phead *ListNode) {
  fmt.Printf("List:")
  for p := phead.Next; p != nil; p = p.Next {
    fmt.Printf("%d,", p.Value)
  }
  fmt.Printf("\n")
}

//末尾添加结点
func AddNodeToTail(phead *ListNode, value int) {
  p := phead
  for p.Next != nil {
    p = p.Next
  }
  p.Next = &ListNode{value, nil}
}

//末尾批量添加结点
func AddNodeListToTail(phead *ListNode, arr1value ...int) {
  p := phead
  for p.Next != nil {
    p = p.Next
  }
  for _, value := range arr1value {
    p.Next = &ListNode{value, nil}
    p = p.Next
  }
}

//删除目标值结点
func DelNodeByValue(phead *ListNode, value int) {
  p := phead
  for p.Next != nil {
    if p.Next.Value == value {
      p.Next = p.Next.Next
    } else {
      p = p.Next
    }
  }
}

//删除倒数第n个结点，代码思路：双指针。
//两个指针都从头开始，一个指针先遍历n个结点，然后另外一个指针开始同步遍历。
//当先遍历的指针移动到链表末尾时，另外一个指针就位于倒数第n个结点。
func DelNthFromTail(phead *ListNode, nth int) {
  var pquery, pprevious *ListNode = phead, phead //遍历指针，倒数第n+1个结点

  if pquery.Next == nil {
    return
  }

  for i := 1; i <= nth+1; i++ { //要删除倒数第n个结点，需要倒数第n+1个结点
    pquery = pquery.Next
    if pquery == nil { //如果链表长度不够倒数，则删除第一个结点
      pprevious.Next = pprevious.Next.Next
      return
    }
  }

  for pquery != nil { //同步遍历
    pquery = pquery.Next
    pprevious = pprevious.Next
  }

  pprevious.Next = pprevious.Next.Next //删除倒数第n个结点
}
