package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
}

//以数组intervals表示若干个区间的集合，其中单个区间为intervals[i]=[starti,endi]
//合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间
//1<=intervals.length<=10^4;intervals[i].length=2;0<=starti<=endi<=10^4

//56-合并区间
func merge(intervals [][]int) [][]int {
	//贪心算法
	//先把所有的区间以左端点升序排序，然后依次和并区间即可
	//如果后一个区间的左端点在前一个区间里，那就可以合并

	var isli2Res [][]int = [][]int{}
	sort.Slice(intervals, func(ii, ij int) bool { return intervals[ii][0] < intervals[ij][0] })

	tl, tr := intervals[0][0], intervals[0][1]
	for ii := 1; ii < len(intervals); ii++ {
		if intervals[ii][0] <= tr {
			if intervals[ii][1] > tr {
				tr = intervals[ii][1]
			}
			continue
		}
		isli2Res = append(isli2Res, []int{tl, tr})
		tl, tr = intervals[ii][0], intervals[ii][1]
	}
	isli2Res = append(isli2Res, []int{tl, tr})

	return isli2Res
}
