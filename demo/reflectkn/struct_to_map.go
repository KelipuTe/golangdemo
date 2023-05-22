package reflectkn

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

// orm 支持的结构体字段的 tag 上的 key 都放在这里
const (
	// 支持的 key 的数量
	tagNum      int    = 1
	tagKeyField string = "fieldWant"
)

func NewErrInputOnlyPointerOrStruct() error {
	return errors.New("orm: 只支持一级结构体指针或者结构体作为输入\n")
}

func NewErrInvalidTagContent(tag string) error {
	return fmt.Errorf("orm: 标签 [%s] 格式错误\n", tag)
}

func F4StructToMap(input any) (map[string]any, error) {
	if nil == input {
		return nil, NewErrInputOnlyPointerOrStruct()
	}

	i9InputType := reflect.TypeOf(input)
	s6InputValue := reflect.ValueOf(input)

	// 只接受一级结构体指针或者结构体
	if reflect.Ptr == i9InputType.Kind() {
		i9InputType = i9InputType.Elem()
		s6InputValue = s6InputValue.Elem()
	}
	if reflect.Struct != i9InputType.Kind() {
		return nil, NewErrInputOnlyPointerOrStruct()
	}

	// 获取结构体字段数量
	fieldNum := i9InputType.NumField()
	m3field := make(map[string]any, fieldNum)
	// 解析结构体的每个字段
	for i := 0; i < fieldNum; i++ {
		s6FieldType := i9InputType.Field(i)
		s6FieldValue := s6InputValue.Field(i)
		m3FieldTag, err := f4ParseTag(s6FieldType.Tag)
		if nil != err {
			return nil, err
		}
		// 从标签里获取设置的数据库字段名
		fieldName := m3FieldTag[tagKeyField]
		// 如果没有设置数据库字段名，默认用转换成小驼峰的结构体字段名
		if "" == fieldName {
			fieldName = f4CamelCaseToSnakeCase(s6FieldType.Name)
		}
		// 私有字段这里是拿不到值的，默认赋 0 值
		if s6FieldType.IsExported() {
			m3field[fieldName] = s6FieldValue.Interface()
		} else {
			m3field[fieldName] = reflect.Zero(s6FieldType.Type).Interface()
		}
	}
	return m3field, nil
}

// f4ParseTag 解析结构体字段的标签
// 标签格式：`orm:"key1=value1,key2=value2"`
func f4ParseTag(s6tag reflect.StructTag) (map[string]string, error) {
	// 从 tag 里面拿 orm 标签
	orm := s6tag.Get("orm")
	if "" == orm {
		return map[string]string{}, nil
	}
	// 解析 tag，其实就是解析字符串
	s5kv := strings.Split(orm, ",")
	m3tag := make(map[string]string, tagNum)
	for _, kv := range s5kv {
		t4s5kv := strings.Split(kv, "=")
		// 判断标签格式正不正确
		if 2 != len(t4s5kv) {
			return nil, NewErrInvalidTagContent(kv)
		}
		m3tag[t4s5kv[0]] = t4s5kv[1]
	}
	return m3tag, nil
}

// f4ToUnderscore 驼峰转蛇形
func f4CamelCaseToSnakeCase(oldString string) string {
	var s5NewString []byte
	for i, char := range oldString {
		// 如果是大写字母，前面加一个下划线，然后转小写字母
		if unicode.IsUpper(char) {
			// 如果首字母是大写字母，不用加下划线
			if 0 != i {
				s5NewString = append(s5NewString, '_')
			}
			s5NewString = append(s5NewString, byte(unicode.ToLower(char)))
		} else {
			s5NewString = append(s5NewString, byte(char))
		}
	}
	return string(s5NewString)
}
