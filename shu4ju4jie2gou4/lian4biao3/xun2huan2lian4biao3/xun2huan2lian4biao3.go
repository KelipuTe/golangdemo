//循环链表
package main

import "fmt"

type ListNode struct {
  Value int
  Next  *ListNode
}

func main() {
  var phead, ptail *ListNode = nil, nil //头指针，尾指针

  tian1jia1duo1ge4(&phead, &ptail, 1, 1, 2, 3, 4, 5, 8, 8)
  shu1chu1(phead)
  shan4chu2by0value(&phead, &ptail, 1)
  shu1chu1(phead)
  shan4chu2by0value(&phead, &ptail, 3)
  shu1chu1(phead)
  shan4chu2by0value(&phead, &ptail, 8)
  shu1chu1(phead)
  tian1jia1duo1ge4(&phead, &ptail, 6, 7)
  shu1chu1(phead)
}

//输出链表
func shu1chu1(phead *ListNode) {
  fmt.Printf("List:")
  if phead != nil {
    fmt.Printf("%d,", phead.Value) //输出头结点
    for pquery := phead.Next; pquery != nil; {
      fmt.Printf("%d,", pquery.Value)
      pquery = pquery.Next
      if pquery == phead {
        break
      }
    }
  }
  fmt.Printf("\n")
}

//末尾添加结点
func tian1jia1yi1ge4(pphead, pptail **ListNode, value int) {
  var pnew *ListNode = &ListNode{value, nil}

  if *pphead == nil { //链表为空
    *pphead = pnew
    *pptail = pnew
    (*pptail).Next = *pphead
  } else { //链表不为空，直接在尾部添加
    (*pptail).Next = pnew
    (*pptail) = pnew
    (*pptail).Next = *pphead
  }
}

//末尾批量添加结点
func tian1jia1duo1ge4(pphead, pptail **ListNode, arr1value ...int) {
  for _, value := range arr1value {
    tian1jia1yi1ge4(pphead, pptail, value)
  }
}

//删除目标值结点
func shan4chu2by0value(pphead, pptail **ListNode, value int) {
  var phead, ptail *ListNode = *pphead, *pptail

  if phead == nil { //链表为空
    return
  }
  if phead == ptail { //只有一个结点
    phead, ptail = nil, nil
    return
  }
  for phead.Value == value { //连续删除头结点
    *pphead = phead.Next
    phead = *pphead
    (*pptail).Next = *pphead //尾结点重接
  }
  for pquery := phead; pquery != nil; { //连续删除头结点之外的结点
    if pquery.Next.Value == value {
      if pquery.Next == ptail { //删除尾结点
        *pptail = pquery
      }
      pquery.Next = pquery.Next.Next
    } else {
      pquery = pquery.Next
    }
    if pquery == phead {
      break
    }
  }
}
