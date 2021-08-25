package main

import "fmt"

func main() {
  fmt.Println(firstBadVersion(5))
}

//通过调用bool isBadVersion(version)接口来判断版本号version是否在单元测试中出错
//假设你有n个版本[1, 2, ..., n]，实现一个函数来查找第一个错误的版本
//1<=bad<=n<=(2^31)-1

//二分查找
//注意，和二分查找不一样，二分查找是数组，是从0开始到n-1，这里是方法调用，是从1开始到n

//278-第一个错误的版本
func firstBadVersion(n int) int {
  indexLeft, indexRight := 1, n
  for indexLeft < indexRight {
    indexMid := indexLeft + (indexRight-indexLeft)>>1
    if isBadVersion(indexMid) {
      //如果中间位置是错的，那么第一个错的就在左边
      indexRight = indexMid
    } else {
      //如果中间位置是对的，那么第一个错的就在右边
      indexLeft = indexMid + 1
    }
  }

  return indexLeft
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool {
  barrRes := []bool{false, false, false, true, true}

  return barrRes[version-1]
}
