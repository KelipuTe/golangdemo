package orm_result

import (
	"database/sql"
	orm_metadata "demo-golang/orm/orm-metadata"
	"reflect"
	"unsafe"
)

// s6ResultUseUnsafe 用 unsafe 实现 I9Result
type s6ResultUseUnsafe struct {
	// p7pointer 存储数据库返回的查询结果的结构体的起始地址
	p7pointer unsafe.Pointer
	// p7s6OrmModel orm 映射模型
	p7s6OrmModel *orm_metadata.S6OrmModel
}

func (p7this s6ResultUseUnsafe) F8SetColumn(rows *sql.Rows) error {
	// 返回数据库字段
	s5ColumnName, err := rows.Columns()
	if nil != err {
		return err
	}
	if len(s5ColumnName) > len(p7this.p7s6OrmModel.M3FieldToStruct) {
		return ErrTooManyReturnedColumns
	}

	s5ColumnValue := make([]interface{}, len(s5ColumnName))
	for i, t4ColumnName := range s5ColumnName {
		// 通过数据库字段找到对应的结构体字段
		p7s6ModelField, ok := p7this.p7s6OrmModel.M3FieldToStruct[t4ColumnName]
		if !ok {
			return F8NewErrUnknownColumn(t4ColumnName)
		}
		// 通过结构体字段的内存偏移量，找到结构体字段的位置
		t4p7pointer := unsafe.Pointer(uintptr(p7this.p7pointer) + p7s6ModelField.Offset)
		// 在找到的找到结构体字段的位置上构造结构体字段
		t4value := reflect.NewAt(p7s6ModelField.I9Type, t4p7pointer)
		s5ColumnValue[i] = t4value.Interface()
	}
	// 从数据库返回的查询结果里取数据
	if err = rows.Scan(s5ColumnValue...); err != nil {
		return err
	}
	return nil
}

// 确保 F8NewS6ResultUseUnsafe 实现的是 F8NewI9Result
var _ F8NewI9Result = F8NewS6ResultUseUnsafe

// F8NewS6ResultUseUnsafe 构造 s6ResultUseUnsafe
func F8NewS6ResultUseUnsafe(value interface{}, p7s5OrmModel *orm_metadata.S6OrmModel) I9Result {
	return &s6ResultUseUnsafe{
		p7pointer:    unsafe.Pointer(reflect.ValueOf(value).Pointer()),
		p7s6OrmModel: p7s5OrmModel,
	}
}
