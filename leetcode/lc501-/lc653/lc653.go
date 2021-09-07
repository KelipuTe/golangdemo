//lc653-两数之和IV-输入BST
//[二叉树][二叉搜索树][中序遍历][数组][双指针][二分查找]

//给定一个二叉搜索树root和一个目标结果k，如果BST中存在两个元素且它们的和等于给定的目标结果，则返回true。
//二叉树的节点个数的范围是[1,10^4];-10^4<=Node.val<=10^4;root为二叉搜索树;-10^5<=k<=10^5

//二叉搜索树的中序遍历是升序序列
//遍历出所有的结点之后，就是有序数组两数之和的问题

package main

import "fmt"

type TreeNode struct {
  Val   int
  Left  *TreeNode
  Right *TreeNode
}

func main() {
  // phead :=
  //   &TreeNode{50,
  //     &TreeNode{30,
  //       &TreeNode{20, nil, nil},
  //       &TreeNode{40, nil, nil}},
  //     &TreeNode{70,
  //       &TreeNode{60, nil, nil},
  //       &TreeNode{80, nil, nil}},
  //   }

  // fmt.Println(findTarget(phead, 60))
  // fmt.Println(findTarget(phead, 20))

  phead :=
    &TreeNode{0,
      &TreeNode{-2,
        nil,
        &TreeNode{-1, nil, nil}},
      &TreeNode{3,
        nil,
        &TreeNode{4, nil, nil}},
    }

  fmt.Println(findTarget(phead, -2))
}

func findTarget(root *TreeNode, k int) bool {
  var sli1Nums []int = []int{}
  var f0zhong1xu4 func(root *TreeNode)

  f0zhong1xu4 = func(root *TreeNode) {
    if root == nil {
      return
    }
    if root.Left != nil {
      f0zhong1xu4(root.Left)
    }
    sli1Nums = append(sli1Nums, root.Val)
    if root.Right != nil {
      f0zhong1xu4(root.Right)
    }
  }

  f0zhong1xu4(root)

  var sli1NumsLen int = len(sli1Nums)

  for index := 0; index < sli1NumsLen; index++ {
    indexLeft, indexRight := index+1, sli1NumsLen-1
    for indexLeft <= indexRight {
      indexMid := indexLeft + (indexRight-indexLeft)>>1
      if sli1Nums[indexMid] > k-sli1Nums[index] {
        indexRight = indexMid - 1
      } else if sli1Nums[indexMid] < k-sli1Nums[index] {
        indexLeft = indexMid + 1
      } else {
        return true
      }
    }
  }

  return false
}
