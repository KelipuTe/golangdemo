package main

import (
	"fmt"
)

func main() {
	name := 1
	fmt.Println(name)
}

//import (
//	"fmt"
//	"unsafe"
//)
//
//func main() {
//	f1 := func() { fmt.Println(1) }
//	f2 := func() { fmt.Println(2) }
//	*(*uintptr)(unsafe.Pointer(&f1)) = *(*uintptr)(unsafe.Pointer(&f2))
//	f1()
//}

//import (
//	"errors"
//	"fmt"
//	"reflect"
//	"strings"
//	"unicode"
//)
//
//func main() {
//	map1, _ := F4StructToMap(S6StructWithTag{Id: 1, Name: "aa", Age: 11, Sex: 1})
//	fmt.Println(map1)
//
//	map2, _ := F4StructToMap(&S6StructWithTag{Id: 2, Name: "bb", Age: 22, Sex: 2})
//	fmt.Println(map2)
//}
//
//type S6StructWithTag struct {
//	Id   int    `orm:"field=id"`
//	Name string `orm:"field=name"`
//	Age  int8   `orm:"field=age"`
//	Sex  int8   `orm:"field=sex"`
//}
//
//// orm 支持的结构体字段的 tag 上的 key 都放在这里
//const (
//	// 支持的 key 的数量
//	tagNum      int    = 1
//	tagKeyField string = "field"
//)
//
//func F4StructToMap(p7s6Ors6 any) (map[string]any, error) {
//	i9msType := reflect.TypeOf(p7s6Ors6)
//	s6msValue := reflect.ValueOf(p7s6Ors6)
//	// 只接受一级结构体指针或者结构体
//	if (reflect.Ptr == i9msType.Kind() && reflect.Struct != i9msType.Elem().Kind()) &&
//		reflect.Struct != i9msType.Kind() {
//		return nil, NewErrInputOnlyPointerOrStruct()
//	}
//	if reflect.Ptr == i9msType.Kind() {
//		i9msType = i9msType.Elem()
//		s6msValue = s6msValue.Elem()
//	}
//
//	// 获取结构体字段数量
//	fieldNum := i9msType.NumField()
//	m3field := make(map[string]any, fieldNum)
//	// 解析结构体的每个字段
//	for i := 0; i < fieldNum; i++ {
//		// 拿字段
//		s6fieldType := i9msType.Field(i)
//		s6fieldValue := s6msValue.Field(i)
//		// 拿字段的 tag
//		m3tag, err := f4ParseTag(s6fieldType.Tag)
//		if nil != err {
//			return nil, err
//		}
//		// 从标签里获取设置的数据库字段名
//		fieldName := m3tag[tagKeyField]
//		// 如果没有设置数据库字段名，默认用转换成小驼峰的结构体字段名
//		if "" == fieldName {
//			fieldName = f4CamelCaseToSnakeCase(s6fieldType.Name)
//		}
//		// 私有字段这里是拿不到值的
//		if s6fieldType.IsExported() {
//			m3field[fieldName] = s6fieldValue.Interface()
//		} else {
//			m3field[fieldName] = reflect.Zero(s6fieldType.Type).Interface()
//		}
//	}
//	return m3field, nil
//}
//
//// f4ParseTag 解析结构体字段的标签
//// 标签格式：`orm:"key1=value1,key2=value2"`
//func f4ParseTag(s6tag reflect.StructTag) (map[string]string, error) {
//	// 从 tag 里面拿 orm 标签
//	orm := s6tag.Get("orm")
//	if "" == orm {
//		return map[string]string{}, nil
//	}
//	// 解析 tag，其实就是解析字符串
//	s5kv := strings.Split(orm, ",")
//	m3tag := make(map[string]string, tagNum)
//	for _, kv := range s5kv {
//		t4s5kv := strings.Split(kv, "=")
//		// 判断标签格式正不正确
//		if 2 != len(t4s5kv) {
//			return nil, NewErrInvalidTagContent(kv)
//		}
//		m3tag[t4s5kv[0]] = t4s5kv[1]
//	}
//	return m3tag, nil
//}
//
//// f4ToUnderscore 驼峰转蛇形
//func f4CamelCaseToSnakeCase(oldString string) string {
//	var s5NewString []byte
//	for i, char := range oldString {
//		// 如果是大写字母，前面加一个下划线，然后转小写字母
//		if unicode.IsUpper(char) {
//			// 如果首字母是大写字母，不用加下划线
//			if 0 != i {
//				s5NewString = append(s5NewString, '_')
//			}
//			s5NewString = append(s5NewString, byte(unicode.ToLower(char)))
//		} else {
//			s5NewString = append(s5NewString, byte(char))
//		}
//	}
//	return string(s5NewString)
//}
//
//func NewErrInputOnlyPointerOrStruct() error {
//	return errors.New("orm: 只支持一级结构体指针或者结构体作为输入\r\n")
//}
//
//func NewErrInvalidTagContent(tag string) error {
//	return fmt.Errorf("orm: 标签 [%s] 格式错误\r\n", tag)
//}

//func main() {
//	//var a interface{} = 1
//	//b, ok := a.(string)
//	//fmt.Println(a, b, ok)
//
//	var a time.Time
//	fmt.Println(a.IsZero())
//}
