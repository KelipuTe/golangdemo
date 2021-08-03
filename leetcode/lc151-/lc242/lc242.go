package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}

//若s和t中每个字符出现的次数都相同，则称s和t互为字母异位词
//给定两个字符串s和t，判断t是否是s的字母异位词
//1<=s.length, t.length<=5*10^4;s和t仅包含小写字母

//242-有效的字母异位词(49,242)(242,438)
func isAnagram(s string, t string) bool {
	//将每个字符串按字符排序，字母异位词排序后的结果一定是一样的

	tbslis := []byte(s)
	sort.Slice(tbslis, func(i, j int) bool { return tbslis[i] < tbslis[j] })
	ts := string(tbslis)
	tbslit := []byte(t)
	sort.Slice(tbslit, func(i, j int) bool { return tbslit[i] < tbslit[j] })
	tt := string(tbslit)
	return ts == tt
}
