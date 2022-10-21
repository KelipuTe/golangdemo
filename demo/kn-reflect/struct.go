package kn_reflect

import (
	"reflect"
)

// IterateStructField 通过反射遍历结构体的字段
func IterateStructField(input any) (map[string]any, error) {
	if nil == input {
		return nil, ErrMustStruct
	}

	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)

	// 处理结构体指针（一级或多级指针）
	for reflect.Pointer == inputType.Kind() {
		inputType = inputType.Elem()
		inputValue = inputValue.Elem()
	}

	if inputType.Kind() != reflect.Struct {
		return nil, ErrMustStruct
	}

	// 获取结构体字段数量
	t1num := inputType.NumField()
	mapres := make(map[string]any, t1num)
	// 解析结构体的每个字段
	for i := 0; i < t1num; i++ {
		t1Field := inputType.Field(i)
		t1FieldVal := inputValue.Field(i)
		// 私有字段这里是拿不到值的，默认赋 0 值
		if t1Field.IsExported() {
			mapres[t1Field.Name] = t1FieldVal.Interface()
		} else {
			mapres[t1Field.Name] = reflect.Zero(t1Field.Type).Interface()
		}
	}

	return mapres, nil
}

// SetStructField 通过反射修改结构体的字段
func SetStructField(instance any, field string, val any) error {
	if nil == instance {
		return ErrMustStructPointer
	}

	inputType := reflect.TypeOf(instance)
	inputValue := reflect.ValueOf(instance)

	// 这里必须是一级结构体指针，其他的都不行
	if inputType.Kind() != reflect.Pointer || inputType.Elem().Kind() != reflect.Struct {
		return ErrMustStructPointer
	}

	inputType = inputType.Elem()
	inputValue = inputValue.Elem()

	// 判断字段存不存在
	if _, ok := inputType.FieldByName(field); !ok {
		return ErrFieldNotFound
	}
	t1FieldVal := inputValue.FieldByName(field)
	// 判断字段能不能赋值
	if !t1FieldVal.CanSet() {
		return ErrFieldCannotSet
	}

	t1FieldVal.Set(reflect.ValueOf(val))

	return nil
}
