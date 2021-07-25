package main

import "fmt"

func main() {
	sNum := "22"
	sarrRes := letterCombinations(sNum)
	fmt.Println(sarrRes)
}

//电话号码的字母组合
//手机九宫格键盘，电话按键2-9分别对应几个字母
//给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合
func letterCombinations(digits string) []string {
	var mapPhone map[string]string = map[string]string{
		"2": "abc", "3": "def", "4": "ghi", "5": "jkl",
		"6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz",
	} //键位和字母表
	var ssli1Res []string  //返回结果
	var tssli1Res []string //保存上一次的结果

	if len(digits) < 1 {
		return []string{}
	}

	//初始化第一个按键的结果
	for ij := 0; ij < len(mapPhone[string(digits[0])]); ij++ {
		var sBtn string = mapPhone[string(digits[0])]
		var cItem byte = sBtn[ij]
		var sRes string = string(cItem)
		ssli1Res = append(ssli1Res, sRes)
	}

	//上一次迭代的结果，依次附加下一个按键对应的所有字母
	for ii := 1; ii < len(digits); ii++ {
		var sBtn string = mapPhone[string(digits[ii])]
		//复制上一次循环的结果
		tssli1Res = make([]string, len(ssli1Res))
		copy(tssli1Res, ssli1Res)
		//清空返回结果
		ssli1Res = make([]string, 0)
		//把这个键位上的字符，依次附加到上一次循环的结果的尾部
		for ij := 0; ij < len(sBtn); ij++ {
			var cItem byte = sBtn[ij]
			for ik := 0; ik < len(tssli1Res); ik++ {
				var sRes string = tssli1Res[ik] + string(cItem)
				ssli1Res = append(ssli1Res, sRes)
			}
		}
	}

	return ssli1Res
}
