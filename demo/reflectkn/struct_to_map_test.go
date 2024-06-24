package reflectkn

//
//import (
//	"fmt"
//	"reflect"
//	"strings"
//	"testing"
//	"unicode"
//)
//
////结构体转map
////结构体（属性）转map（map[属性名]属性值）
////结构体（属性+tag）转map（map[tag]属性值）
//
//func structToMapByTag(in any) (map[string]any, error) {
//	if in == nil {
//		return nil, ErrMustStructOrPointer
//	}
//
//	inrt := reflect.TypeOf(in)
//	inrv := reflect.ValueOf(in)
//
//	if inrt.Kind() != reflect.Pointer {
//		return nil, ErrMustStructPointer
//	}
//	inrt = inrt.Elem()
//	inrv = inrv.Elem()
//	if inrt.Kind() != reflect.Struct {
//		return nil, ErrMustStructPointer
//	}
//
//	fieldNum := inrt.NumField()
//	fieldMap := make(map[string]any, fieldNum)
//	for i := 0; i < fieldNum; i++ {
//		fieldrt := inrt.Field(i)
//		fieldrv := inrv.Field(i)
//		tagMap, err2 := parseTag(fieldrt.Tag)
//		if err2 != nil {
//			return nil, err2
//		}
//		// 从标签里获取设置的数据库字段名
//		fieldName := tagMap["field"]
//		// 如果没有设置数据库字段名，默认用转换成小驼峰的结构体字段名
//		if "" == fieldName {
//			fieldName = f4CamelCaseToSnakeCase(fieldrt.Name)
//		}
//		// 私有字段这里是拿不到值的，默认赋 0 值
//		if fieldrt.IsExported() {
//			fieldMap[fieldName] = fieldrv.Interface()
//		} else {
//			fieldMap[fieldName] = reflect.Zero(fieldrt.Type).Interface()
//		}
//	}
//	return fieldMap, nil
//}
//
//// parseTag 解析结构体字段的标签
//// 这里的格式是：`orm:"field=name"`
//func parseTag(tag reflect.StructTag) (map[string]string, error) {
//	// 从 tag 里面拿 orm 标签
//	ormStr := tag.Get("orm")
//	if ormStr == "" {
//		return map[string]string{}, nil
//	}
//
//	// 解析 tag，其实就是解析字符串
//	ormSplit := strings.Split(ormStr, ",")
//	ormMap := make(map[string]string, 1)
//	for _, v := range ormSplit {
//		vSplit := strings.Split(v, "=")
//		// 判断标签格式正不正确
//		if len(vSplit) != 2 {
//			return nil, fmt.Errorf("标签 [%s] 格式错误\n", v)
//		}
//		ormMap[vSplit[0]] = vSplit[1]
//	}
//	return ormMap, nil
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
//type S6StructWithTag struct {
//	Id   int    `orm:"wantField=id"`
//	Name string `orm:"wantField=name"`
//	Age  int8   `orm:"wantField=privateV1"`
//	Sex  int8   `orm:"wantField=sex"`
//}
//
//func TestStructToMap(p7tt *testing.T) {
//	s5case := []struct {
//		name    string
//		input   any
//		mapWant map[string]any
//		errWant error
//	}{
//		{
//			name: "struct",
//			input: S6StructWithTag{
//				Id:   1,
//				Name: "aa",
//				Age:  11,
//				Sex:  1,
//			},
//			mapWant: map[string]any{
//				"id":        1,
//				"name":      "aa",
//				"privateV1": int8(11),
//				"sex":       int8(1),
//			},
//			errWant: nil,
//		},
//		{
//			name: "pointer",
//			input: &S6StructWithTag{
//				Id:   1,
//				Name: "aa",
//				Age:  11,
//				Sex:  1,
//			},
//			mapWant: map[string]any{
//				"id":        1,
//				"name":      "aa",
//				"privateV1": int8(11),
//				"sex":       int8(1),
//			},
//			errWant: nil,
//		},
//	}
//
//	for _, t4case := range s5case {
//		p7tt.Run(t4case.name, func(p7tt *testing.T) {
//			m3res, err := structToMap(t4case.input)
//			assert.Equal(p7tt, t4case.errWant, err)
//			if err != nil {
//				return
//			}
//			assert.Equal(p7tt, t4case.mapWant, m3res)
//		})
//	}
//}
