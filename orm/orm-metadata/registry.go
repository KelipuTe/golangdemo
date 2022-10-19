package orm_metadata

import (
	"reflect"
	"strings"
	"sync"
	"unicode"
)

// F4OrmModelOption Option 设计模式
type F4OrmModelOption func(p7s6om *S6OrmModel) error

// I9Registry 对元数据注册中心的抽象
type I9Registry interface {
	// F4Get 查找元数据
	F4Get(p7s6Model any) (*S6OrmModel, error)
	// F4Register 注册元数据
	F4Register(p7s6Model any, s5f4Option ...F4OrmModelOption) (*S6OrmModel, error)
}

// s6Registry 元数据注册中心
type s6Registry struct {
	// m3Model 解析好的 orm 映射模型
	// 这里可以预计会遇到并发操作，所以用 sync.Map
	m3Model sync.Map
}

func (p7this *s6Registry) F4Get(p7s6Model any) (*S6OrmModel, error) {
	i9type := reflect.TypeOf(p7s6Model)
	// 如果结构体已经解析过了，就直接返回
	value, ok := p7this.m3Model.Load(i9type)
	if ok {
		return value.(*S6OrmModel), nil
	}
	// 否则需要解析并注册新的 orm 映射模型
	return p7this.F4Register(p7s6Model)
}

func (p7this *s6Registry) F4Register(p7s6Model any, s5f4Option ...F4OrmModelOption) (*S6OrmModel, error) {
	p7s6om, err := p7this.f4ParseModel(p7s6Model)
	if nil != err {
		return nil, err
	}
	for _, t4f4 := range s5f4Option {
		err = t4f4(p7s6om)
		if nil != err {
			return nil, err
		}
	}
	i9type := reflect.TypeOf(p7s6Model)
	p7this.m3Model.Store(i9type, p7s6om)
	return p7s6om, nil
}

// f4ParseModel 解析结构体
func (p7this *s6Registry) f4ParseModel(p7s6Model any) (*S6OrmModel, error) {
	i9msType := reflect.TypeOf(p7s6Model)
	// 只接受一级结构体指针
	if reflect.Ptr != i9msType.Kind() || reflect.Struct != i9msType.Elem().Kind() {
		return nil, NewErrInputOnlyPointer()
	}
	i9msType = i9msType.Elem()

	// 获取表名
	var tableName string
	i9tn, ok := p7s6Model.(I9TableName)
	if ok {
		tableName = i9tn.F4TableName()
	}
	if "" == tableName {
		tableName = f4CamelCaseToSnakeCase(i9msType.Name())
	}

	// 获取结构体字段数量
	fieldNum := i9msType.NumField()
	m3stf := make(map[string]*S6ModelField, fieldNum)
	m3fts := make(map[string]*S6ModelField, fieldNum)
	// 解析结构体的每个字段
	for i := 0; i < fieldNum; i++ {
		s6fieldType := i9msType.Field(i)
		m3tag, err := p7this.f4ParseTag(s6fieldType.Tag)
		if nil != err {
			return nil, err
		}
		// 从标签里获取设置的数据库字段名
		fieldName := m3tag[tagKeyField]
		// 如果没有设置数据库字段名，默认用转换成小驼峰的结构体字段名
		if "" == fieldName {
			fieldName = f4CamelCaseToSnakeCase(s6fieldType.Name)
		}
		// 正反方向都要存一份
		p7s6mf := &S6ModelField{
			StructName: s6fieldType.Name,
			I9Type:     s6fieldType.Type,
			Offset:     s6fieldType.Offset,
			FieldName:  fieldName,
		}
		m3stf[s6fieldType.Name] = p7s6mf
		m3fts[fieldName] = p7s6mf
	}

	p7om := &S6OrmModel{
		TableName:       tableName,
		M3StructToField: m3stf,
		M3FieldToStruct: m3fts,
	}

	return p7om, nil
}

// f4ParseTag 解析结构体字段的标签
// 标签格式：`orm:"key1=value1,key2=value2"`
func (p7this *s6Registry) f4ParseTag(s6tag reflect.StructTag) (map[string]string, error) {
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

func NewI9Registry() I9Registry {
	return &s6Registry{}
}
