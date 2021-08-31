package main

import "fmt"

func main() {
  fmt.Println(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5))
}

//这里有n个航班，它们分别从1到n进行编号。
//有一份航班预订表bookings，表中第i条预订记录bookings[i]=[firsti,lasti,seatsi]
//意味着在从firsti到lasti（包含firsti和lasti）的每个航班上预订了seatsi个座位。
//请你返回一个长度为n的数组answer，其中answer[i]是航班i上预订的座位总数。

//数组，差分
//差分数组的第i个数即为原数组的第i−1个元素和第i个元素的差值。数组[1,2,2,4]，其差分数组为[1,1,0,2]。
//差分数组对应的概念是前缀和数组，也就是说对差分数组求前缀和即可得到原数组。
//当对原数组的某一个区间[l,r]施加一个增量a时，差分数组d的d[l]增加a，d[r+1]减少a。
//这样对于区间的修改就变为了对于两个位置的修改，并且这种修改是可以叠加的。
//举例，对区间[1, 2]增加10，对区间[2, 3]增加20，对区间[2, 5]增加25。
//原数组  [0,0,0,0,0]=>[10,10,  0,0,0]=>[10,30, 20,  0,0]=>[10,55, 45, 25,25]
//差分数组[0,0,0,0,0]=>[10, 0,-10,0,0]=>[10,20,-10,-20,0]=>[10,45,-10,-20, 0]
//这里只有5个位置，如果后面还有第六个位置的话。差分数组应该是这样的，[10,45,-10,-20, 0,-25]

//1109-航班预订统计
func corpFlightBookings(bookings [][]int, n int) []int {
  var sli1Res []int = make([]int, n)

  for _, item := range bookings {
    sli1Res[item[0]-1] += item[2]
    if item[1] < n {
      sli1Res[item[1]] -= item[2]
    }
  }

  for index := 1; index < n; index++ {
    sli1Res[index] += sli1Res[index-1]
  }

  return sli1Res
}
