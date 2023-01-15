package main

import "fmt"

// 链式处理
func main() {
	arr1 := []int{1, 2, 3, 4, 5}

	// 需要一组输入和输出相同的方法
	chain := []func(int) int{
		add1,
		times2,
	}

	fmt.Printf("arr1: %+v\n", arr1)

	for i, v := range arr1 {
		t1result := v
		// 依次调用所有的方法
		// 前一个方法的输出作为后一个方法的输入
		for _, t1func := range chain {
			t1result = t1func(t1result)
		}
		arr1[i] = t1result
	}

	fmt.Printf("arr1: %+v\n", arr1)
}

func add1(t1 int) int {
	return t1 + 1
}

func times2(t1 int) int {
	return t1 << 1
}
