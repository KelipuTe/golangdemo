package main

import (
	"fmt"
	"sort"
)

func main() {
	// combinationSum([]int{2, 3, 6, 7}, 7)
	combinationSum([]int{2, 3, 5}, 8)
	// combinationSum([]int{1}, 1)
	fmt.Println(isli2Res)
}

//给定一个无重复元素的正整数数组candidates和一个正整数target
//找出candidates中所有可以使数字和为目标数target的唯一组合
//candidates中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的
//1<=candidates.length<=30;1<=candidates[i]<=200;1<=target<=500

var isli2Res [][]int //结果集

//39-组合总和(39,46,47)
func combinationSum(candidates []int, target int) [][]int {
	//回溯算法
	//因为可以重复，所以每次回溯都有两种选择，要自己或者跳过自己
	isli2Res = [][]int{}         //初始化结果集
	var isli1Res []int = []int{} //每次回溯的结果
	sort.Ints(candidates)        //排序，有利于回溯时直接过滤掉超过target的元素
	hui2Su4(candidates, target, 0, isli1Res)
	return isli2Res
}

func hui2Su4(candidates []int, target int, iIndex int, isli1Res []int) {
	if target == 0 {
		//目标数为0证明已经找到一个组合
		var tsli1Res []int = make([]int, len(isli1Res))
		copy(tsli1Res, isli1Res)
		isli2Res = append(isli2Res, tsli1Res)
		return
	} else if target < 0 {
		return
	}

	//要自己
	isli1Res = append(isli1Res, candidates[iIndex])
	hui2Su4(candidates, target-candidates[iIndex], iIndex, isli1Res)
	isli1Res = isli1Res[:len(isli1Res)-1]
	if iIndex+1 < len(candidates) {
		//跳过自己，直接下一个
		hui2Su4(candidates, target, iIndex+1, isli1Res)
	}
}
