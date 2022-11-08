package metadata

import (
	"reflect"
	"strings"
	"sync"
	"unicode"
)

// F8S6OrmModelOption 方法抽象：Option 设计模式
type F8S6OrmModelOption func(p7s6om *S6OrmModel) error

// I9Registry 接口抽象：对元数据注册中心的抽象
type I9Registry interface {
	// F8Get 方法抽象：查找元数据
	F8Get(p7s6Model any) (*S6OrmModel, error)
	// F8Register 方法抽象：注册元数据
	F8Register(p7s6Model any, s5f8Option ...F8S6OrmModelOption) (*S6OrmModel, error)
}

// s6Registry 元数据注册中心
// 实现 I9Registry 接口
type s6Registry struct {
	// m3Model map：解析好的 orm 映射模型
	// 这里可以预见会出现并发操作，所以用 sync.Map
	m3Model sync.Map
}

func (p7this *s6Registry) F8Get(p7s6Model any) (*S6OrmModel, error) {
	// 看看传进来的结构体解析过没有
	i9type := reflect.TypeOf(p7s6Model)
	value, ok := p7this.m3Model.Load(i9type.String())
	// 如果结构体已经解析过了，就直接返回
	if ok {
		return value.(*S6OrmModel), nil
	}
	// 否则需要解析并注册 orm 映射模型
	return p7this.F8Register(p7s6Model)
}

func (p7this *s6Registry) F8Register(p7s6Model any, s5f8Option ...F8S6OrmModelOption) (*S6OrmModel, error) {
	p7s6om, err := p7this.f8ParseModel(p7s6Model)
	if nil != err {
		return nil, err
	}
	// 执行 Option 设计模式
	for _, t4f8 := range s5f8Option {
		err = t4f8(p7s6om)
		if nil != err {
			return nil, err
		}
	}
	// 注册 orm 映射模型
	i9type := reflect.TypeOf(p7s6Model)
	p7this.m3Model.Store(i9type.String(), p7s6om)
	return p7s6om, nil
}

// f8ParseModel 解析结构体
func (p7this *s6Registry) f8ParseModel(p7s6Model any) (*S6OrmModel, error) {
	i9ModelType := reflect.TypeOf(p7s6Model)
	// 只接受一级结构体指针
	if reflect.Ptr != i9ModelType.Kind() || reflect.Struct != i9ModelType.Elem().Kind() {
		return nil, F8NewErrInputOnlyStructPointer()
	}
	i9ModelType = i9ModelType.Elem()

	// 获取表名
	var tableName string
	t4i9TableName, ok := p7s6Model.(I9TableName)
	if ok {
		tableName = t4i9TableName.F8TableName()
	}
	if "" == tableName {
		tableName = f8CamelCaseToSnakeCase(i9ModelType.Name())
	}

	// 获取结构体字段数量
	fieldNum := i9ModelType.NumField()
	m3SToF := make(map[string]*S6ModelField, fieldNum)
	m3FToS := make(map[string]*S6ModelField, fieldNum)
	// 解析结构体的每个字段
	for i := 0; i < fieldNum; i++ {
		s6FieldType := i9ModelType.Field(i)
		m3tag, err := p7this.f8ParseTag(s6FieldType.Tag)
		if nil != err {
			return nil, err
		}
		// 从标签里获取设置的数据库字段名
		fieldName := m3tag[tagKeyField]
		// 如果没有设置数据库字段名，默认用转换成小驼峰的结构体字段名
		if "" == fieldName {
			fieldName = f8CamelCaseToSnakeCase(s6FieldType.Name)
		}
		// 正反方向都要存一份
		p7s6mf := &S6ModelField{
			StructName: s6FieldType.Name,
			I9Type:     s6FieldType.Type,
			Offset:     s6FieldType.Offset,
			FieldName:  fieldName,
		}
		m3SToF[s6FieldType.Name] = p7s6mf
		m3FToS[fieldName] = p7s6mf
	}

	p7s6om := &S6OrmModel{
		TableName:       tableName,
		M3StructToField: m3SToF,
		M3FieldToStruct: m3FToS,
	}

	return p7s6om, nil
}

// f8ParseTag 解析结构体字段的标签
// 标签格式：`orm:"key1=value1,key2=value2"`
func (p7this *s6Registry) f8ParseTag(s6tag reflect.StructTag) (map[string]string, error) {
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
			return nil, F8NewErrInvalidTagContent(kv)
		}
		m3tag[t4s5kv[0]] = t4s5kv[1]
	}
	return m3tag, nil
}

// f8ToUnderscore 驼峰转蛇形
func f8CamelCaseToSnakeCase(oldString string) string {
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

// F8NewI9Registry 构造 I9Registry
// s6Registry 是 I9Registry 的实例
func F8NewI9Registry() I9Registry {
	return &s6Registry{}
}
