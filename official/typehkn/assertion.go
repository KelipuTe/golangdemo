package typehkn

import "fmt"

type testStruct struct {
	value int
}

type testStructV2 struct {
	value string
}

// 类型断言
func f8TypeAssertion(input any) {
	_, ok := input.(int)
	if ok {
		fmt.Println("int")
	} else {
		fmt.Println("not int")
	}
}

// 类型断言
func f8TypeAssertionV2(input any) {
	switch input.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

// 类型强转
func f8ForcedConversion(input any) {
	value, ok := input.(testStruct)
	if ok {
		fmt.Println("testStruct.value:", value.value)
	} else {
		fmt.Println("not testStruct")
	}
}

// 类型强转
func f8ForcedConversionV2(input any) {
	// 这么写， value 就是断言过的变量，case 里面可以直接当已知类型用
	switch value := input.(type) {
	case testStruct:
		fmt.Println("testStruct.value:", value.value)
	case testStructV2:
		fmt.Println("testStructV2.value:", value.value)
	default:
		fmt.Println("unknown")
	}
}
