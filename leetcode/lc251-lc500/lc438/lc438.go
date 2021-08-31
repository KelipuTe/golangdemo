package main

import (
	"fmt"
)

func main() {
	// fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}

//异位词指字母相同，但排列不同的字符串
//给定两个字符串s和p，找到s中所有p的异位词的子串，返回这些子串的起始索引

//438-找到字符串中所有字母异位词(242,438)
func findAnagrams(s string, p string) []int {
	//滑动窗口
	//用两个长度为26的数组，分别存储滑动窗口和目标字符串中每个字母出现的次数

	var ssli1Res []int = []int{}
	var iarr1S, iarr1P [26]int = [26]int{}, [26]int{}
	var bsli1S []byte = []byte(s)
	var bsli1P []byte = []byte(p)

	iPLen := len(bsli1P)
	for ii := 0; ii < iPLen; ii++ {
		iarr1P[bsli1P[ii]-'a']++
	}

	iSLen := len(bsli1S)
	for ii, ij := 0, 0; ij < iSLen; {
		if ij-ii <= iPLen-1 {
			iarr1S[bsli1S[ij]-'a']++
			ij++
		} else {
			if iarr1S == iarr1P {
				ssli1Res = append(ssli1Res, ii)
			}
			iarr1S[bsli1S[ii]-'a']--
			ii++
		}
		if ij == iSLen {
			//判断最后依次滑动
			if iarr1S == iarr1P {
				ssli1Res = append(ssli1Res, ii)
			}
		}
	}

	return ssli1Res
}
