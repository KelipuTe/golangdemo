package reflect

import (
	"errors"
	"reflect"
)

var ErrMustStructOrStructPointer = errors.New("input must be struct or struct pointer")

type StructFuncInfo struct {
	Name      string
	SliInput  []reflect.Type
	SliOutput []reflect.Type
	SliRes    []any
}

// IterateStructFunc 通过反射遍历结构体的方法
func IterateStructFunc(input any) (map[string]*StructFuncInfo, error) {
	t1type := reflect.TypeOf(input)
	if t1type.Kind() != reflect.Struct && t1type.Kind() != reflect.Pointer {
		return nil, ErrMustStructOrStructPointer
	}

	// 有几个方法
	t1Num := t1type.NumMethod()
	mapres := make(map[string]*StructFuncInfo, t1Num)
	// 遍历方法
	for i := 0; i < t1Num; i++ {
		t1func := t1type.Method(i)

		// 有几个输入参数
		t1InputNum := t1func.Type.NumIn()
		t1sliInput := make([]reflect.Type, 0, t1InputNum)
		// 构造调用方法需要的输入参数，第一个参数永远都是接收器
		t1sliCallInput := make([]reflect.Value, 0, t1InputNum)
		t1sliCallInput = append(t1sliCallInput, reflect.ValueOf(input))

		// 遍历输入参数
		for j := 0; j < t1InputNum; j++ {
			t1param := t1func.Type.In(j)
			t1sliInput = append(t1sliInput, t1param)

			if j > 0 {
				t1sliCallInput = append(t1sliCallInput, reflect.Zero(t1param))
			}
		}

		// 调用方法
		t1sliCallOutput := t1func.Func.Call(t1sliCallInput)

		// 有几个输出参数
		t1OutputNun := t1func.Type.NumOut()
		t1sliOutput := make([]reflect.Type, 0, t1OutputNun)
		// 记录调用方法输出的结果
		t1slires := make([]any, 0, t1OutputNun)

		// 遍历输出参数
		for j := 0; j < t1OutputNun; j++ {
			t1sliOutput = append(t1sliOutput, t1func.Type.Out(j))
			t1slires = append(t1slires, t1sliCallOutput[j].Interface())
		}

		mapres[t1func.Name] = &StructFuncInfo{
			Name:      t1func.Name,
			SliInput:  t1sliInput,
			SliOutput: t1sliOutput,
			SliRes:    t1slires,
		}
	}

	return mapres, nil
}
