//lc232-用栈实现队列
//[数组][栈][队列]

//仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）
//实现MyQueue类：
//voidpush(int x)将元素x推到队列的末尾。intpop()从队列的开头移除并返回元素。
//intpeek()返回队列开头的元素。booleanempty()如果队列为空，返回true；否则，返回false。
//说明：
//只能使用标准的栈操作——也就是只有pushtotop,peek/popfromtop,size,和isempty操作是合法的。
//使用的语言也许不支持栈。可以使用list或者deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
//进阶：
//能否实现每个操作均摊时间复杂度为O(1)的队列？
//换句话说，执行n个操作的总时间复杂度为O(n)，即使其中一个操作可能花费较长时间。

//1<=x<=9；最多调用100次push、pop、peek和empty；
//假设所有操作都是有效的（例如，一个空的队列不会调用pop或者peek操作）

//两个栈，入队栈用于入队，出队栈用于出队。
//入队的时候，数据入栈入队栈，出队的时候，先检查出队栈是否为空。
//如果出队栈为空，就依次出栈入队栈的数据，然后把数据入栈出队栈。
//如果出队栈不为空，直接出栈。

package main

import "fmt"

func main() {
  queue := Constructor()
  queue.Push(1)
  queue.Push(2)
  queue.Push(3)
  queue.Push(4)
  fmt.Println(queue.Pop())
  queue.Push(5)
  fmt.Println(queue.Pop())
  fmt.Println(queue.Pop())
  fmt.Println(queue.Pop())
  fmt.Println(queue.Pop())
}

type MyQueue struct {
  stackIn, stackOut       [100]int
  stackInLen, stackOutLen int
}

func Constructor() MyQueue {
  return MyQueue{[100]int{}, [100]int{}, 0, 0}
}

func (this *MyQueue) Push(x int) {
  this.stackIn[this.stackInLen] = x
  this.stackInLen++
}

func (this *MyQueue) Pop() int {
  if this.stackOutLen == 0 {
    for index := this.stackInLen - 1; index > -1; index-- {
      this.stackOut[this.stackOutLen] = this.stackIn[index]
      this.stackOutLen++
    }
    this.stackInLen = 0
  }
  this.stackOutLen--
  return this.stackOut[this.stackOutLen]
}

func (this *MyQueue) Peek() int {
  if this.stackOutLen == 0 {
    return this.stackIn[0]
  }
  return this.stackOut[this.stackOutLen-1]
}

func (this *MyQueue) Empty() bool {
  return this.stackOutLen == 0 && this.stackInLen == 0
}
