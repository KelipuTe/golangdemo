package main

import (
  "fmt"
  "math"
)

func main() {
  var minStack MinStack = Constructor()
  minStack.Push(-2)
  minStack.Push(0)
  minStack.Push(-3)
  fmt.Println(minStack.GetMin())
  minStack.Pop()
  fmt.Println(minStack.Top())
  fmt.Println(minStack.GetMin())
}

//设计一个支持push，pop，top操作，并能在常数时间内检索到最小元素的栈。
//push(x)——将元素x推入栈中。
//pop()——删除栈顶的元素。
//top()——获取栈顶元素。
//getMin()——检索栈中的最小元素。

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

//对于栈中每一个元素来说，最小值是不一样的，所以不能简单地用一个变量来存储最小值
//这里使用一个辅助栈来存储每个栈元素对应的最小值

//155-最小栈

type MinStack struct {
  iSize         int   //栈大小
  isli1Zhan4    []int //栈
  isli1MinZhan4 []int //辅助栈，用于存储最小值
}

/** initialize your data structure here. */
func Constructor() MinStack {
  var ms MinStack = MinStack{
    iSize:         0,
    isli1Zhan4:    []int{},
    isli1MinZhan4: []int{},
  }
  return ms
}

func (this *MinStack) Push(val int) {
  this.isli1Zhan4 = append(this.isli1Zhan4, val)
  iMinTop := math.MaxInt32 //如果栈是空栈，就默认给最大临界值
  if this.iSize > 0 {
    iMinTop = this.isli1MinZhan4[this.iSize-1]
  }
  if val < iMinTop {
    this.isli1MinZhan4 = append(this.isli1MinZhan4, val)
  } else {
    this.isli1MinZhan4 = append(this.isli1MinZhan4, iMinTop)
  }
  this.iSize++
}

func (this *MinStack) Pop() {
  this.isli1Zhan4 = this.isli1Zhan4[:this.iSize-1]
  this.isli1MinZhan4 = this.isli1MinZhan4[:this.iSize-1]
  this.iSize--
}

func (this *MinStack) Top() int {
  return this.isli1Zhan4[this.iSize-1]
}

func (this *MinStack) GetMin() int {
  return this.isli1MinZhan4[this.iSize-1]
}
