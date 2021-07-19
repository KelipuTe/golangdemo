package main

import "fmt"

func main() {
	sNum := "22"
	sarrRes := letterCombinations(sNum)
	fmt.Println(sarrRes)
}

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
	var ssliRes []string  //返回结果
	var tssliRes []string //保存上一次的结果

	if len(digits) < 1 {
		return []string{}
	}

	//初始化第一个
	for ij := 0; ij < len(mapPhone[string(digits[0])]); ij++ {
		sBtn := mapPhone[string(digits[0])]
		cItem := sBtn[ij]
		sRes := string(cItem)
		ssliRes = append(ssliRes, sRes)
	}
	//处理后面的
	for ii := 1; ii < len(digits); ii++ {
		sBtn := mapPhone[string(digits[ii])]
		//复制上一次循环的结果
		tssliRes = make([]string, len(ssliRes))
		copy(tssliRes, ssliRes)
		//清空返回结果
		ssliRes = make([]string, 0)
		//把这个键位上的字符，依次附加到上一次循环的结果的尾部
		for ij := 0; ij < len(sBtn); ij++ {
			cItem := sBtn[ij]
			for ik := 0; ik < len(tssliRes); ik++ {
				sRes := tssliRes[ik] + string(cItem)
				ssliRes = append(ssliRes, sRes)
			}
		}
	}

	return ssliRes
}
