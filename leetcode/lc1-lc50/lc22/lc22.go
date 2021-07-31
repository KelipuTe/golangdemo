package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

//数字n代表生成括号的对数，生成所有有效的括号组合

//22-括号生成(17,22)
func generateParenthesis(n int) []string {
	//动态规划
	//对于n对括号，可以分解为1对括号和n-1对括号的位置的问题
	//固定一对括号后，剩下的n-1对，可能在这对括号里面或者右边
	//左边也可以，但是如果左边也算上，就会重复得出一些结论
	//假设括号里面a对括号右边b对，a+b=n-1
	//那么n对括号的结果，就变成了："("+a对的结果+")"+b对的结果
	//1对：(0对)0对=>"("")"；结果：()
	//2对：(1对)0对=>"("()")"；(0对)1对=>"("")"()；结果：(()),()()
	//3对：(2对)0对；(1对)1对；(0对)2对

	var mapRes map[int][]string = map[int][]string{0: {}, 1: {"()"}}

	if n < 1 {
		return mapRes[0]
	} else if n == 1 {
		return mapRes[1]
	}

	for iRes := 2; iRes <= n; iRes++ {
		var ssliRes []string = []string{}
		//注意，里层循环的控制变量是变化的
		for iZuo3 := 0; iZuo3 < iRes; iZuo3++ {
			iYou4 := iRes - 1 - iZuo3
			isli1Zuo3Len := len(mapRes[iZuo3])
			isli1You4Len := len(mapRes[iYou4])
			if isli1Zuo3Len == 0 {
				//(0对)n-1对
				tsZuo3 := ""
				for iiYou4 := 0; iiYou4 < isli1You4Len; iiYou4++ {
					tsYou4 := ""
					if isli1You4Len != 0 {
						tsYou4 = mapRes[iYou4][iiYou4]
					}
					tsRes := "(" + tsZuo3 + ")" + tsYou4
					ssliRes = append(ssliRes, tsRes)
				}
			} else {
				//(a对)b对
				for iiZuo3 := 0; iiZuo3 < isli1Zuo3Len; iiZuo3++ {
					tsZuo3 := mapRes[iZuo3][iiZuo3]
					if isli1You4Len == 0 {
						//(n-1对)0对
						tsYou4 := ""
						tsRes := "(" + tsZuo3 + ")" + tsYou4
						ssliRes = append(ssliRes, tsRes)
					} else {
						//(a对)b对
						for iiYou4 := 0; iiYou4 < isli1You4Len; iiYou4++ {
							tsYou4 := mapRes[iYou4][iiYou4]
							tsRes := "(" + tsZuo3 + ")" + tsYou4
							ssliRes = append(ssliRes, tsRes)
						}
					}
				}
			}
			mapRes[iRes] = ssliRes
		}
	}

	return mapRes[n]
}
