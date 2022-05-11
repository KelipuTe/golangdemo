package main

import "fmt"

// 剑指 Offer 09. 用两个栈实现队列
// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
// 分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

// 解题思路
// 一个栈用于入队；一个栈用于出队
// 入队的时候，直接加入入队栈即可；出队的时候，如果出队栈没有数据，就把入队栈的数据依次出栈然后加入出队栈。

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
func main() {
  obj := Constructor()
  fmt.Println(obj.DeleteHead())
  obj.AppendTail(5)
  obj.AppendTail(2)
  fmt.Println(obj.DeleteHead())
  fmt.Println(obj.DeleteHead())
  fmt.Println(obj.DeleteHead())
  obj.AppendTail(1)
  obj.AppendTail(8)
  obj.AppendTail(20)
  obj.AppendTail(1)
  obj.AppendTail(11)
  obj.AppendTail(2)
  fmt.Println(obj.DeleteHead())
  fmt.Println(obj.DeleteHead())
  fmt.Println(obj.DeleteHead())
  fmt.Println(obj.DeleteHead())
}

type CQueue struct {
  sli1In  []int // 入队栈
  sli1Out []int // 出队栈
}

func Constructor() CQueue {
  return CQueue{
    sli1In:  make([]int, 0, 16),
    sli1Out: make([]int, 0, 16),
  }
}

func (this *CQueue) AppendTail(value int) {
  this.sli1In = append(this.sli1In, value)
}

func (this *CQueue) DeleteHead() int {
  sli1OutLen := len(this.sli1Out)
  if 0 < sli1OutLen {
    num := this.sli1Out[sli1OutLen-1]
    this.sli1Out = this.sli1Out[:sli1OutLen-1]
    return num
  } else {
    sli1InLen := len(this.sli1In)
    if 0 >= sli1InLen {
      return -1
    }
    for i := sli1InLen - 1; i > 0; i-- {
      this.sli1Out = append(this.sli1Out, this.sli1In[i])
    }
    num := this.sli1In[0]
    this.sli1In = make([]int, 0, 16)
    return num
  }
}
