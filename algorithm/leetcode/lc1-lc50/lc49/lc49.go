package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{"a"}))
	// fmt.Println(groupAnagrams([]string{""}))
}

//给定一个字符串数组，将字母异位词组合在一起
//字母异位词指字母相同，但排列不同的字符串
//可以按任意顺序返回结果列表

//49-字母异位词分组(49,242)
func groupAnagrams(strs []string) [][]string {
	//将每个字符串按字符排序，字母异位词排序后的结果一定是一样的
	//构造一个map，用这个排序后的结果作为键，将原来的字符串分类

	var ssli2Res [][]string = [][]string{}
	var tmapRes map[string][]string = map[string][]string{}

	for ii := 0; ii < len(strs); ii++ {
		tsValue := []byte(strs[ii])
		sort.Slice(tsValue, func(i, j int) bool { return tsValue[i] < tsValue[j] })
		tsKey := string(tsValue)
		tmapRes[tsKey] = append(tmapRes[tsKey], strs[ii])
	}

	for _, ssli1Value := range tmapRes {
		ssli2Res = append(ssli2Res, ssli1Value)
	}

	return ssli2Res
}
