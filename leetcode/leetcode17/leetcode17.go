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
	mapPhone := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	} //键位和字母表
	var sarr1sliRes []string  //返回结果
	var tsarr1sliRes []string //保存上一次的结果

	if len(digits) < 1 {
		return []string{}
	}

	//初始化第一个
	for ij := 0; ij < len(mapPhone[string(digits[0])]); ij++ {
		sBtn := mapPhone[string(digits[0])]
		cItem := sBtn[ij]
		sRes := string(cItem)
		sarr1sliRes = append(sarr1sliRes, sRes)
	}
	//处理后面的
	for ii := 1; ii < len(digits); ii++ {
		sBtn := mapPhone[string(digits[ii])]
		//复制上一次循环的结果
		tsarr1sliRes = make([]string, len(sarr1sliRes))
		copy(tsarr1sliRes, sarr1sliRes)
		//清空返回结果
		sarr1sliRes = make([]string, 0)
		//把这个键位上的字符，依次附加到上一次循环的结果的尾部
		for ij := 0; ij < len(sBtn); ij++ {
			cItem := sBtn[ij]
			for ik := 0; ik < len(tsarr1sliRes); ik++ {
				sRes := tsarr1sliRes[ik] + string(cItem)
				sarr1sliRes = append(sarr1sliRes, sRes)
			}
		}
	}

	return sarr1sliRes
}
