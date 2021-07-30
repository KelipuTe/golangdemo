package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(5))
}

//通过调用bool isBadVersion(version)接口来判断版本号version是否在单元测试中出错
//假设你有n个版本[1, 2, ..., n]，实现一个函数来查找第一个错误的版本

//278、第一个错误的版本
func firstBadVersion(n int) int {
	//二分查找
	iLeft, iRight := 1, n
	for iLeft < iRight {
		iMid := iLeft + (iRight-iLeft)>>1
		if isBadVersion(iMid) {
			iRight = iMid
		} else {
			iLeft = iMid + 1
		}
	}
	return iLeft
}

func isBadVersion(version int) bool {
	barrRes := []bool{false, false, false, true, true}
	return barrRes[version-1]
}
