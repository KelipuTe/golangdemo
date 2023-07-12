package result

import (
	"database/sql"
	"demo-golang/orm/internal"
	"demo-golang/orm/metadata"
	"reflect"
)

// 确保 F8NewS6ResultUseReflect 实现的是 F8NewI9Result
var _ F8NewI9Result = F8NewS6ResultUseReflect

// s6ResultUseReflect 用反射实现 I9Result
type s6ResultUseReflect struct {
	// s6Value 存储数据库返回的查询结果的结构体
	s6Value reflect.Value
	// p7s6Model 映射模型
	p7s6Model *metadata.S6Model
}

func (p7this s6ResultUseReflect) F8GetField(name string) (any, error) {
	res := p7this.s6Value.FieldByName(name)
	if res == (reflect.Value{}) {
		return nil, internal.F8NewErrUnknownColumn(name)
	}
	return res.Interface(), nil
}

func (p7this s6ResultUseReflect) F8SetField(rows *sql.Rows) error {
	// 返回数据库列名
	s5ColumnName, err := rows.Columns()
	if nil != err {
		return err
	}
	if len(s5ColumnName) > len(p7this.p7s6Model.M3ColumnToField) {
		return internal.ErrTooManyReturnedColumns
	}

	// 这里初始化的时候需要长度
	// s5ColumnValue 和 s5ColumnValueElem 最终都指向同一个对象
	s5ColumnValue := make([]any, len(s5ColumnName))
	s5ColumnValueElem := make([]reflect.Value, len(s5ColumnName))
	for i, t4ColumnName := range s5ColumnName {
		// 通过数据库列名找到对应的结构体属性
		p7s6ModelField, ok := p7this.p7s6Model.M3ColumnToField[t4ColumnName]
		if !ok {
			return internal.F8NewErrUnknownColumn(t4ColumnName)
		}
		// 构造结构体属性
		t4value := reflect.New(p7s6ModelField.I9Type)
		s5ColumnValue[i] = t4value.Interface()
		s5ColumnValueElem[i] = t4value.Elem()
	}
	// 从数据库返回的查询结果里取数据
	if err = rows.Scan(s5ColumnValue...); err != nil {
		return err
	}
	for i, t4ColumnName := range s5ColumnName {
		// 通过数据库列名找到对应的结构体属性
		p7s6ModelField := p7this.p7s6Model.M3ColumnToField[t4ColumnName]
		t4value := p7this.s6Value.FieldByName(p7s6ModelField.FieldName)
		// 把取到的数据放到结构体属性上
		t4value.Set(s5ColumnValueElem[i])
	}
	return nil
}

// F8NewS6ResultUseReflect s6ResultUseReflect 的构造方法
func F8NewS6ResultUseReflect(value any, p7s5OrmModel *metadata.S6Model) I9Result {
	return &s6ResultUseReflect{
		s6Value:   reflect.ValueOf(value).Elem(),
		p7s6Model: p7s5OrmModel,
	}
}
