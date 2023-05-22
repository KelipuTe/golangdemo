package reflectkn

import (
	"reflect"
)

// f8IterateStructField 通过反射遍历结构体的字段
func f8IterateStructField(input any) (map[string]any, error) {
	if nil == input {
		return nil, ErrMustStructOrStructPointer
	}

	i9InputType := reflect.TypeOf(input)
	s6InputValue := reflect.ValueOf(input)

	// 处理结构体指针（一级或多级指针）
	for reflect.Pointer == i9InputType.Kind() {
		i9InputType = i9InputType.Elem()
		s6InputValue = s6InputValue.Elem()
	}
	if reflect.Struct != i9InputType.Kind() {
		return nil, ErrMustStructOrStructPointer
	}

	// 结构体字段数量
	fieldNum := i9InputType.NumField()
	m3result := make(map[string]any, fieldNum)
	// 解析结构体的每个字段
	for i := 0; i < fieldNum; i++ {
		s6FieldType := i9InputType.Field(i)
		s6FieldValue := s6InputValue.Field(i)
		// 私有字段这里是拿不到值的，默认赋零值
		if s6FieldType.IsExported() {
			m3result[s6FieldType.Name] = s6FieldValue.Interface()
		} else {
			m3result[s6FieldType.Name] = reflect.Zero(s6FieldType.Type).Interface()
		}
	}

	return m3result, nil
}

// f8SetStructField 通过反射修改结构体的字段
func f8SetStructField(input any, field string, value any) error {
	// 因为需要修改结构体，所以必须是一级结构体指针
	if nil == input {
		return ErrMustStructPointer
	}

	i9InputType := reflect.TypeOf(input)
	s6InputValue := reflect.ValueOf(input)

	// 处理结构体指针
	if reflect.Pointer == i9InputType.Kind() {
		i9InputType = i9InputType.Elem()
		s6InputValue = s6InputValue.Elem()
	}
	if reflect.Struct != i9InputType.Kind() {
		return ErrMustStructPointer
	}

	// 判断字段存不存在
	if _, ok := i9InputType.FieldByName(field); !ok {
		return ErrStructFieldNotFound
	}
	s6FieldValue := s6InputValue.FieldByName(field)
	// 判断字段能不能赋值
	if !s6FieldValue.CanSet() {
		return ErrStructFieldCannotSet
	}
	s6FieldValue.Set(reflect.ValueOf(value))

	return nil
}

type S6FuncInfo struct {
	// 方法名
	Name string
	// 方法的入参的类型
	S5InputType []reflect.Type
	// 方法的出参的类型
	S5OutputType []reflect.Type
	// 方法的出参的值
	S5OutputValue []any
}

// f8IterateStructFunc 通过反射遍历结构体的方法
func f8IterateStructFunc(input any) (map[string]*S6FuncInfo, error) {
	if nil == input {
		return nil, ErrMustStructOrStructPointer
	}

	i9InputType := reflect.TypeOf(input)

	// 处理结构体指针
	if reflect.Pointer == i9InputType.Kind() && reflect.Struct != i9InputType.Elem().Kind() {
		return nil, ErrMustStructOrStructPointer
	}
	if reflect.Struct != i9InputType.Kind() {
		return nil, ErrMustStructOrStructPointer
	}

	// 结构体方法数量
	funcNum := i9InputType.NumMethod()
	m3result := make(map[string]*S6FuncInfo, funcNum)
	// 解析结构体的每个方法
	for i := 0; i < funcNum; i++ {
		t4func := i9InputType.Method(i)
		// 方法的入参的数量
		funcInputNum := t4func.Type.NumIn()
		s5FuncInput := make([]reflect.Type, 0, funcInputNum)
		// 方法的出参的数量
		funcOutputNum := t4func.Type.NumOut()
		s5FuncOutput := make([]reflect.Type, 0, funcOutputNum)
		s5res := make([]any, 0, funcOutputNum)

		// 构造调用方法需要的入参
		s5FuncCallInput := make([]reflect.Value, 0, funcInputNum)
		// 注意，第一个参数永远都是接收器
		s5FuncCallInput = append(s5FuncCallInput, reflect.ValueOf(input))
		// 按下标遍历入参
		for j := 0; j < funcInputNum; j++ {
			// 反射得到入参的类型
			funcInputType := t4func.Type.In(j)
			s5FuncInput = append(s5FuncInput, funcInputType)
			// 用入参的类型的 0 值构造请求参数
			if j > 0 {
				s5FuncCallInput = append(s5FuncCallInput, reflect.Zero(funcInputType))
			}
		}

		// 用上面构造的请求参数调用方法
		s5FuncCallOutput := t4func.Func.Call(s5FuncCallInput)

		// 按下标遍历出参
		for j := 0; j < funcOutputNum; j++ {
			// 反射得到出参的类型
			funcOutputType := t4func.Type.Out(j)
			s5FuncOutput = append(s5FuncOutput, funcOutputType)
			// 记录调用方法后得到的出参的值
			s5res = append(s5res, s5FuncCallOutput[j].Interface())
		}

		m3result[t4func.Name] = &S6FuncInfo{
			Name:          t4func.Name,
			S5InputType:   s5FuncInput,
			S5OutputType:  s5FuncOutput,
			S5OutputValue: s5res,
		}
	}

	return m3result, nil
}
