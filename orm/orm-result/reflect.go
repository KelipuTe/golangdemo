package orm_result

import (
	"database/sql"
	orm_metadata "demo-golang/orm/orm-metadata"
	"reflect"
)

// 用反射实现
type resultUseReflect struct {
	//
	s6value      reflect.Value
	p7s6OrmModel *orm_metadata.S6OrmModel
}

func (p7this resultUseReflect) F8SetColumn(rows *sql.Rows) error {
	cs, err := rows.Columns()
	if err != nil {
		return err
	}

	// colValues 和 colEleValues 实质上最终都指向同一个对象
	colValues := make([]interface{}, len(cs))
	colEleValues := make([]reflect.Value, len(cs))
	for i, c := range cs {
		cm, ok := p7this.p7s6OrmModel.M3FieldToStruct[c]
		if !ok {
			return NewErrUnknownColumn(c)
		}
		val := reflect.New(cm.I9Type)
		colValues[i] = val.Interface()
		colEleValues[i] = val.Elem()
	}
	if err = rows.Scan(colValues...); err != nil {
		return err
	}
	for i, c := range cs {
		cm := p7this.p7s6OrmModel.M3FieldToStruct[c]
		fd := p7this.s6value.FieldByName(cm.StructName)
		fd.Set(colEleValues[i])
	}
	return nil
}

var _ F8NewI9Result = F8NewResultUseReflect

func F8NewResultUseReflect(value interface{}, p7s5OrmModel *orm_metadata.S6OrmModel) I9Result {
	return &resultUseReflect{
		s6value:      reflect.ValueOf(value).Elem(),
		p7s6OrmModel: p7s5OrmModel,
	}
}
