package orm

import (
	"context"
	"database/sql"
	"demo-golang/orm/internal"
)

// S6UpdateBuilder UPDATE 查询构造器
type S6UpdateBuilder[T any] struct {
	// p7Entity 需要修改的实体，解析它得到元数据
	p7Entity *T
	// s5i9Assignment SET 后面的
	s5i9Assignment []i9Assignment
	// s5where WHERE 后面的
	s5where []S6WhereCondition

	i9Session I9Session
	s6QueryBuilder
}

func (p7this *S6UpdateBuilder[T]) F8SetEntity(p7Entity *T) *S6UpdateBuilder[T] {
	if nil == p7Entity {
		return p7this
	}
	p7this.p7Entity = p7Entity
	return p7this
}

func (p7this *S6UpdateBuilder[T]) F8SetUpdate(s5i9Assignment ...i9Assignment) *S6UpdateBuilder[T] {
	if 0 >= len(s5i9Assignment) {
		return p7this
	}
	if nil == p7this.s5i9Assignment {
		p7this.s5i9Assignment = s5i9Assignment
		return p7this
	}
	p7this.s5i9Assignment = append(p7this.s5i9Assignment, s5i9Assignment...)
	return p7this
}

func (p7this *S6UpdateBuilder[T]) F8Where(s5condition ...S6WhereCondition) *S6UpdateBuilder[T] {
	if 0 >= len(s5condition) {
		return p7this
	}
	if nil == p7this.s5where {
		p7this.s5where = s5condition
		return p7this
	}
	p7this.s5where = append(p7this.s5where, s5condition...)
	return p7this
}

func (p7this *S6UpdateBuilder[T]) F8BuildQuery() (*S6Query, error) {
	var err error = nil

	if nil == p7this.p7Entity {
		p7this.p7Entity = new(T)
	}
	if 0 >= len(p7this.s5i9Assignment) {
		return nil, internal.ErrUpdateWithoutColumn
	}
	if 0 >= len(p7this.s5where) {
		return nil, internal.ErrUpdateWithoutWhere
	}

	p7this.s6QueryBuilder.p7s6Model, err = p7this.s6Monitor.i9Registry.F8Get(p7this.p7Entity)
	if nil != err {
		return nil, err
	}

	p7this.s6QueryBuilder.sqlString.WriteString("UPDATE ")
	p7this.s6QueryBuilder.f8WrapWithQuote(p7this.s6QueryBuilder.p7s6Model.TableName)
	p7this.s6QueryBuilder.sqlString.WriteString(" SET ")

	// 通过反射拿到结构体属性的值
	t4i9result := p7this.f8NewI9Result(p7this.p7Entity, p7this.p7s6Model)
	for i, t4value := range p7this.s5i9Assignment {
		if 0 < i {
			p7this.s6QueryBuilder.sqlString.WriteByte(',')
		}
		switch t4value2 := t4value.(type) {
		case S6Column:
			p7s6ModelField, ok := p7this.s6QueryBuilder.p7s6Model.M3FieldToColumn[t4value2.fieldName]
			if !ok {
				return nil, internal.F8NewErrUnknownField(t4value2.fieldName)
			}
			p7this.s6QueryBuilder.f8WrapWithQuote(p7s6ModelField.ColumnName)
			p7this.s6QueryBuilder.sqlString.WriteString("=?")
			t4EntityValue, err2 := t4i9result.F8GetField(p7s6ModelField.FieldName)
			if nil != err2 {
				return nil, err2
			}
			p7this.f8AddParameter(t4EntityValue)
		case S6Assignment:
			err = t4value2.f8BuildAssignment(&p7this.s6QueryBuilder)
			if nil != err {
				return nil, err
			}
		}
	}

	p7this.sqlString.WriteString(" WHERE ")
	err = p7this.f8BuildWhereCondition(p7this.s5where)
	if nil != err {
		return nil, err
	}

	p7this.s6QueryBuilder.sqlString.WriteByte(';')

	return &S6Query{
		SQLString: p7this.s6QueryBuilder.sqlString.String(),
		S5Value:   p7this.s6QueryBuilder.s5Value,
	}, nil
}

func (p7this *S6UpdateBuilder[T]) F8EXEC(ctx context.Context) (sql.Result, error) {
	p7s6Context := &S6QueryContext{
		QueryType: "UPDATE",
		i9Builder: p7this,
		p7s6Model: p7this.s6QueryBuilder.p7s6Model,
		p7s6Query: nil,
	}
	p7s6Result := f8DoEXEC(ctx, p7this.i9Session, &p7this.s6Monitor, p7s6Context)
	return p7s6Result.I9SQLResult, p7s6Result.I9Err
}

func F8NewS6UpdateBuilder[T any](i9Session I9Session) *S6UpdateBuilder[T] {
	t4p7s6monitor := i9Session.f8GetS6Monitor()
	return &S6UpdateBuilder[T]{
		i9Session: i9Session,
		s6QueryBuilder: s6QueryBuilder{
			s6Monitor: t4p7s6monitor,
			quote:     t4p7s6monitor.i9Dialect.f8GetQuoter(),
		},
	}
}
