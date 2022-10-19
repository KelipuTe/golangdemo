package orm_metadata

import "reflect"

// I9TableName 实现这个接口来返回自定义的表名
type I9TableName interface {
	F4TableName() string
}

// S6OrmModel orm 映射模型
// 处理结构体字段和数据库字段的互相转换
type S6OrmModel struct {
	// TableName 结构体对应的表名
	TableName string
	// M3StructToField 结构体字段 => 数据库字段
	M3StructToField map[string]*S6ModelField
	// M3FieldToStruct 数据库字段 => 结构体字段
	M3FieldToStruct map[string]*S6ModelField
}

// S6ModelField orm 映射模型的每个字段
type S6ModelField struct {
	// 结构体字段名
	StructName string
	// I9Type 结构体字段类型
	I9Type reflect.Type
	// 结构体字段相对于对象的起始地址的偏移量
	Offset uintptr
	// 数据库字段名
	FieldName string
}

// orm 支持的结构体字段的 tag 上的 key 都放在这里
const (
	// 支持的 key 的数量
	tagNum      int    = 1
	tagKeyField string = "field"
)
