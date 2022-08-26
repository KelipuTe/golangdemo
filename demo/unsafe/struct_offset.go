package unsafe

import (
	"fmt"
	"reflect"
)

// PrintStructFieldOffset 输出结构体的内存布局，字段大小和字段起始地址相对于结构体起始地址的偏移量
func PrintStructFieldOffset(entity any) {
	t4type := reflect.TypeOf(entity)
	fmt.Printf("type:%8s, size:%2d\r\n", t4type.Name(), t4type.Size())
	for i := 0; i < t4type.NumField(); i++ {
		t4field := t4type.Field(i)
		fmt.Printf("field:%8s,type:%8s, size:%2d, offset:%2d\r\n", t4field.Name, t4field.Type.Name(), t4field.Type.Size(), t4field.Offset)
	}
	fmt.Printf("\r\n")
}
