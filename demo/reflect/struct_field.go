package reflect

import (
	"reflect"
)

// IterateStructField 通过反射遍历结构体的字段
func IterateStructField(input any) (map[string]any, error) {
	if nil == input {
		return nil, ErrMustStruct
	}

	t1type := reflect.TypeOf(input)
	t1val := reflect.ValueOf(input)

	// 处理结构体指针（一级或多级指针）
	for reflect.Pointer == t1type.Kind() {
		t1type = t1type.Elem()
		t1val = t1val.Elem()
	}

	if t1type.Kind() != reflect.Struct {
		return nil, ErrMustStruct
	}

	// 有几个字段
	t1num := t1type.NumField()
	mapres := make(map[string]any, t1num)
	// 遍历字段
	for i := 0; i < t1num; i++ {
		t1Field := t1type.Field(i)
		t1FieldVal := t1val.Field(i)
		// 私有字段这里是拿不到值的
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

	t1type := reflect.TypeOf(instance)
	t1val := reflect.ValueOf(instance)

	// 这里必须是一级结构体指针，其他的都不行
	if t1type.Kind() != reflect.Pointer || t1type.Elem().Kind() != reflect.Struct {
		return ErrMustStructPointer
	}

	t1type = t1type.Elem()
	t1val = t1val.Elem()

	// 判断字段存不存在
	if _, ok := t1type.FieldByName(field); !ok {
		return ErrFieldNotFound
	}
	t1FieldVal := t1val.FieldByName(field)
	// 判断字段能不能赋值
	if !t1FieldVal.CanSet() {
		return ErrFieldCannotSet
	}

	t1FieldVal.Set(reflect.ValueOf(val))

	return nil
}
