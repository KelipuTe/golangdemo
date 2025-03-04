package main

type Node struct {
  Val   int
  Left  *Node
  Right *Node
  Next  *Node
}

func main() {
  tn7 := Node{7, nil, nil, nil}
  tn6 := Node{6, nil, nil, nil}
  tn5 := Node{5, nil, nil, nil}
  tn4 := Node{4, nil, nil, nil}
  tn3 := Node{3, &tn6, &tn7, nil}
  tn2 := Node{2, &tn4, &tn5, nil}
  tn1 := Node{1, &tn2, &tn3, nil}

  connect(&tn1)
}

//给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
//填充它的每个next指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将next指针设置为NULL。
//初始状态下，所有next指针都被设置为NULL。
//进阶：只能使用常量级额外空间。
//进阶：使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

//二叉树，队列，层次遍历

//116-填充每个节点的下一个右侧节点指针
func connect(root *Node) *Node {
  var queue []*Node = []*Node{}
  var startIndex, endIndex int = 0, 0

  if root == nil {
    return root
  }

  queue = append(queue, root)
  endIndex++

  for startIndex < endIndex {
    endIndexNow := endIndex
    for startIndex < endIndexNow {
      pnow := queue[startIndex]
      if startIndex+1 < endIndexNow {
        //如果后一个结点是这一层里的就连接
        pnow.Next = queue[startIndex+1]
      }
      //添加这个结点的左右结点
      if pnow.Left != nil {
        queue = append(queue, pnow.Left)
        endIndex++
      }
      if pnow.Right != nil {
        queue = append(queue, pnow.Right)
        endIndex++
      }
      startIndex++
    }
  }

  return root
}
