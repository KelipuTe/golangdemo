package result

import (
	"database/sql"
	orm_metadata "demo-golang/orm/metadata"
	"reflect"
)

// s6ResultUseReflect 用反射实现 I9Result
type s6ResultUseReflect struct {
	// s6value 存储数据库返回的查询结果的结构体
	s6value reflect.Value
	// p7s6OrmModel orm 映射模型
	p7s6OrmModel *orm_metadata.S6OrmModel
}

func (p7this s6ResultUseReflect) F8SetColumn(rows *sql.Rows) error {
	// 返回数据库字段
	s5ColumnName, err := rows.Columns()
	if nil != err {
		return err
	}
	if len(s5ColumnName) > len(p7this.p7s6OrmModel.M3FieldToStruct) {
		return ErrTooManyReturnedColumns
	}

	// s5ColumnValue 和 s5ColumnValueElem 最终都指向同一个对象
	s5ColumnValue := make([]interface{}, len(s5ColumnName))
	s5ColumnValueElem := make([]reflect.Value, len(s5ColumnName))
	for i, t4ColumnName := range s5ColumnName {
		// 通过数据库字段找到对应的结构体字段
		p7s6ModelField, ok := p7this.p7s6OrmModel.M3FieldToStruct[t4ColumnName]
		if !ok {
			return F8NewErrUnknownColumn(t4ColumnName)
		}
		// 构造结构体字段
		t4value := reflect.New(p7s6ModelField.I9Type)
		s5ColumnValue[i] = t4value.Interface()
		s5ColumnValueElem[i] = t4value.Elem()
	}
	// 从数据库返回的查询结果里取数据
	if err = rows.Scan(s5ColumnValue...); err != nil {
		return err
	}
	for i, t4ColumnName := range s5ColumnName {
		// 通过数据库字段找到对应的结构体字段
		p7s6ModelField := p7this.p7s6OrmModel.M3FieldToStruct[t4ColumnName]
		t4value := p7this.s6value.FieldByName(p7s6ModelField.StructName)
		// 把取到的数据放到结构体字段上
		t4value.Set(s5ColumnValueElem[i])
	}
	return nil
}

// 确保 F8NewS6ResultUseReflect 实现的是 F8NewI9Result
var _ F8NewI9Result = F8NewS6ResultUseReflect

// F8NewS6ResultUseReflect s6ResultUseReflect 的构造方法
func F8NewS6ResultUseReflect(value interface{}, p7s5OrmModel *orm_metadata.S6OrmModel) I9Result {
	return &s6ResultUseReflect{
		s6value:      reflect.ValueOf(value).Elem(),
		p7s6OrmModel: p7s5OrmModel,
	}
}
