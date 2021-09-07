//lc198-打家劫舍
//[动态规划]

package main

import "fmt"

func main() {
  fmt.Println(rob([]int{0}))
  // fmt.Println(rob([]int{1, 2, 3, 1}))
  // fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

//一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。
//影响偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//给定一个代表每个房屋存放金额的非负整数数组，计算不触动警报装置的情况下，一夜之内能够偷窃到的最高金额。
//1<=nums.length<=100;0<=nums[i]<=400

//举例[1,2,3,1]
//如果只有1间，则只有唯一选择[1]，如果有2间，则选择大的那间[2]
//如果有3间，那么第3间选或不选，取决于，第3间加上第1间的结果[1+3]和前2间的结果[2]哪个大
//如果有4间，那么第4间选或不选，取决于，第4间加上前2间的结果[2+1]和前3间的结果[1+3]哪个大
//由此可以推导出，第n间选或不选，取决于，第n间加上前n-2间的结果和前n-1间的结果哪个大
//第n间选或不选，取决于，第n-1间选或不选，是一个向前连续依赖的关系，有状态转移关系，所以可以用动态规划计算

func rob(nums []int) int {
  var iNumsLen int
  var isli1MaxRes []int

  iNumsLen = len(nums)
  if iNumsLen < 1 {
    return 0
  }
  if iNumsLen < 2 {
    return nums[0]
  }

  isli1MaxRes = make([]int, iNumsLen)
  //初始化前1间和前2间
  isli1MaxRes[0] = nums[0]
  if nums[1] > nums[0] {
    isli1MaxRes[1] = nums[1]
  } else {
    isli1MaxRes[1] = nums[0]
  }
  //计算后面的
  for iIndex := 2; iIndex < iNumsLen; iIndex++ {
    if nums[iIndex]+isli1MaxRes[iIndex-2] > isli1MaxRes[iIndex-1] {
      isli1MaxRes[iIndex] = nums[iIndex] + isli1MaxRes[iIndex-2]
    } else {
      isli1MaxRes[iIndex] = isli1MaxRes[iIndex-1]
    }
  }

  return isli1MaxRes[iNumsLen-1]
}
