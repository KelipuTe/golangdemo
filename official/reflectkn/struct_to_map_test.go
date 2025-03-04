package reflectkn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
	"unicode"
)

//结构体转map
//属性无tag => map[属性名]属性值
//属性有tag => map[tag]属性值

// 结构体转map
func structToMap(in any) (map[string]any, error) {
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
	fieldMap := make(map[string]any, fieldNum)
	for i := 0; i < fieldNum; i++ {
		vrt := irt.Field(i)
		vrv := irv.Field(i)

		// 从标签里获取设置的转换名
		tagMap, err2 := parseTag(vrt.Tag)
		if err2 != nil {
			return nil, err2
		}
		transName := tagMap["name"]

		// 如果没有设置，默认把属性名转小写蛇形作为转换名
		if transName == "" {
			transName = camelToSnake(vrt.Name)
		}

		if vrt.IsExported() {
			fieldMap[transName] = vrv.Interface()
		} else {
			fieldMap[transName] = reflect.Zero(vrt.Type).Interface()
		}
	}
	return fieldMap, nil
}

// parseTag 解析结构体字段的标签
// 这里的格式是：`tag:"name=name"`
func parseTag(tag reflect.StructTag) (map[string]string, error) {
	// 从 tag 里面拿 orm 标签
	tagStr := tag.Get("tag")
	if tagStr == "" {
		return map[string]string{}, nil
	}

	// 解析 tag，其实就是解析字符串
	tagSplit := strings.Split(tagStr, ",")
	tagMap := make(map[string]string, 1)
	for _, v := range tagSplit {
		vSplit := strings.Split(v, "=")
		// 判断标签格式正不正确
		if len(vSplit) != 2 {
			return nil, fmt.Errorf("标签 [%s] 格式错误\n", v)
		}
		tagMap[vSplit[0]] = vSplit[1]
	}

	return tagMap, nil
}

// camelToSnake 驼峰转蛇形
func camelToSnake(oldStr string) string {
	var newStr []byte

	for i, v := range oldStr {
		// 如果是大写字母，前面加一个下划线，然后转小写字母
		if unicode.IsUpper(v) {
			// 如果首字母是大写字母，不用加下划线
			if i != 0 {
				newStr = append(newStr, '_')
			}
			newStr = append(newStr, byte(unicode.ToLower(v)))
		} else {
			newStr = append(newStr, byte(v))
		}
	}

	return string(newStr)
}

type CaseUserV3 struct {
	ID   int `tag:"name=id"`
	Name string
}

func TestStructToMap(t *testing.T) {
	caseList := []struct {
		name    string
		input   any
		wantRes map[string]any
		wantErr error
	}{
		{
			name:    "普通结构体",
			input:   CaseUserV3{ID: 1, Name: "aa"},
			wantRes: map[string]any{"id": 1, "name": "aa"},
			wantErr: nil,
		},
		{
			name:    "一级指针",
			input:   &CaseUserV3{ID: 1, Name: "aa"},
			wantRes: map[string]any{"id": 1, "name": "aa"},
			wantErr: nil,
		},
	}

	for _, v := range caseList {
		t.Run(v.name, func(t *testing.T) {
			res, err := structToMap(v.input)
			assert.Equal(t, v.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, v.wantRes, res)
		})
	}
}
