package main

import "fmt"

func main() {
  fmt.Println(numTrees(3))
}

//给你一个整数n ，求恰由n个节点组成且节点值从1到n互不相同的二叉搜索树有多少种

//动态规划
//n个数，可以选取任意一点作为根结点，假设选第i个数作为根结点，那么问题就变成
//i左边的1到i-1，可以构成多少个二叉搜索树，i右边的i+1到n，可以构成多少个二叉搜索树
//这两个问题的结果的乘积，就是第i个数作为根结点可以构成的二叉搜索树的个数
//然后i从1遍历到n依次累加就是最终的结果
//这里后面的结果是依赖前面的结果的，所以就可以使用动态规划

//96-不同的二叉搜索树
func numTrees(n int) int {
  var isli1Res []int = make([]int, n+1) //结果集

  isli1Res[0] = 1
  //下标0的位置初始化为1是因为，有一边是空树时，总数量就是另一边的数量
  //做乘法的时候，只要不是两边都是空树，空树的情况可以理解为是1种情况
  isli1Res[1] = 1

  for ii := 2; ii <= n; ii++ {
    isli1Res[ii] = 0
    for ij := 1; ij <= ii; ij++ {
      isli1Res[ii] += isli1Res[ij-1] * isli1Res[ii-ij]
    }
  }

  return isli1Res[n]
}
