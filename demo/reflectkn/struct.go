package reflectkn

import (
	"reflect"
)

//通过反射访问结构体的属性和方法
//获取属性信息，修改属性值，获取属性的tag，获取方法信息，访问方法

// 获取属性信息
func visitStructField(in any) (map[string]any, error) {
	if in == nil {
		return nil, ErrMustStructOrPointer
	}

	irt := reflect.TypeOf(in)
	irv := reflect.ValueOf(in)

	// 处理结构体指针（一级或多级指针）
	for irt.Kind() == reflect.Pointer {
		irt = irt.Elem()
		irv = irv.Elem()
	}
	if irt.Kind() != reflect.Struct {
		return nil, ErrMustStructOrPointer
	}

	// 有几个属性，按下标遍历然后解析
	fieldNum := irt.NumField()
	fieldList := make(map[string]any, fieldNum)
	for i := 0; i < fieldNum; i++ {
		// 类型反射和值反射分别获取属性信息
		vrt := irt.Field(i)
		vrv := irv.Field(i)

		// 判断是不是公开属性
		if vrt.IsExported() {
			fieldList[vrt.Name] = vrv.Interface()
		} else {
			// 私有属性是拿不到值的，默认赋零值
			fieldList[vrt.Name] = reflect.Zero(vrt.Type).Interface()
		}
	}

	return fieldList, nil
}

func visitStructTag(in any) (map[string]string, error) {
	if in == nil {
		return nil, ErrMustStructOrPointer
	}

	irt := reflect.TypeOf(in)
	irv := reflect.ValueOf(in)

	for irt.Kind() == reflect.Pointer {
		irt = irt.Elem()
		irv = irv.Elem()
	}
	if irt.Kind() != reflect.Struct {
		return nil, ErrMustStructOrPointer
	}

	fieldNum := irt.NumField()
	fieldList := make(map[string]string, fieldNum)
	for i := 0; i < fieldNum; i++ {
		vrt := irt.Field(i)
		fieldList[vrt.Name] = string(vrt.Tag)
	}

	return fieldList, nil
}

// 修改属性值
func editStructField(in any, field string, value any) error {
	// 因为需要修改结构体，所以必须是一级指针
	if in == nil {
		return ErrMustStructPointer
	}

	irt := reflect.TypeOf(in)
	irv := reflect.ValueOf(in)

	// 处理结构体指针
	if irt.Kind() != reflect.Pointer {
		return ErrMustStructPointer
	}
	irt = irt.Elem()
	irv = irv.Elem()
	if irt.Kind() != reflect.Struct {
		return ErrMustStructPointer
	}

	// 判断属性存不存在
	if _, ok := irt.FieldByName(field); !ok {
		return ErrStructFieldNotFound
	}
	vrv := irv.FieldByName(field)

	// 判断属性能不能赋值
	if !vrv.CanSet() {
		return ErrStructFieldCannotSet
	}
	vrv.Set(reflect.ValueOf(value))

	return nil
}

type FuncInfo struct {
	Name         string         // 方法名
	InTypeList   []reflect.Type // 入参类型
	OutTypeList  []reflect.Type // 出参类型
	OutValueList []any          // 出参值
}

// 获取方法信息
func visitStructFunc(in any) (map[string]*FuncInfo, error) {
	if in == nil {
		return nil, ErrMustStructOrPointer
	}

	irt := reflect.TypeOf(in)

	if irt.Kind() != reflect.Struct && irt.Kind() != reflect.Pointer {
		return nil, ErrMustStructOrPointer
	}
	if irt.Kind() == reflect.Pointer && irt.Elem().Kind() != reflect.Struct {
		return nil, ErrMustStructOrPointer
	}

	// 有几个方法，按下标解析每个方法
	funcNum := irt.NumMethod()
	funcList := make(map[string]*FuncInfo, funcNum)
	for i := 0; i < funcNum; i++ {
		vrt := irt.Method(i)

		// 入参数量，按下标遍历
		inNum := vrt.Type.NumIn()
		inTypeList := make([]reflect.Type, 0, inNum)
		for j := 0; j < inNum; j++ {
			// 反射得到参数类型
			v2 := vrt.Type.In(j)
			inTypeList = append(inTypeList, v2)
		}

		// 出参数量，按下标遍历
		outNum := vrt.Type.NumOut()
		outTypeList := make([]reflect.Type, 0, outNum)
		for j := 0; j < outNum; j++ {
			// 反射得到参数类型
			v2 := vrt.Type.Out(j)
			outTypeList = append(outTypeList, v2)
		}

		funcList[vrt.Name] = &FuncInfo{
			Name:        vrt.Name,
			InTypeList:  inTypeList,
			OutTypeList: outTypeList,
		}
	}

	return funcList, nil
}

// 调用方法
func callStructFunc(in any) (map[string]*FuncInfo, error) {
	funcList, err := visitStructFunc(in)
	if err != nil {
		return nil, err
	}

	irt := reflect.TypeOf(in)
	funcNum := irt.NumMethod()
	for i := 0; i < funcNum; i++ {
		vrt := irt.Method(i)
		v := funcList[vrt.Name]

		//构造调用方法需要的入参。注意，第一个参数永远都是接收器。
		inNum := len(v.InTypeList)
		inValueList := make([]reflect.Value, 0, inNum)
		inValueList = append(inValueList, reflect.ValueOf(in))
		for j := 0; j < inNum; j++ {
			if j < 1 {
				continue
			}
			// 这里意思一下，用入参类型的 0 值构造请求参数
			inValueList = append(inValueList, reflect.Zero(v.InTypeList[j]))
		}

		// 用上面构造的请求参数调用方法
		resValueList := vrt.Func.Call(inValueList)

		// 记录调用方法后得到的出参的值
		outNum := len(v.OutTypeList)
		outValueList := make([]any, 0, outNum)
		for j := 0; j < outNum; j++ {
			outValueList = append(outValueList, resValueList[j].Interface())
		}

		funcList[vrt.Name].OutValueList = outValueList
	}

	return funcList, nil
}
