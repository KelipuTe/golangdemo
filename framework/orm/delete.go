package orm

import (
	"context"
	"database/sql"
	"demo-golang/orm/internal"
)

// S6DeleteBuilder DELETE 查询构造器
type S6DeleteBuilder[T any] struct {
	// p7Entity 需要修改的实体，解析它得到元数据
	p7Entity *T
	// s5where WHERE 后面的
	s5where []S6WhereCondition

	i9Session I9Session
	s6QueryBuilder
}

func F8NewS6DeleteBuilder[T any](i9Session I9Session) *S6DeleteBuilder[T] {
	t4p7s6monitor := i9Session.f8GetS6Monitor()
	return &S6DeleteBuilder[T]{
		i9Session: i9Session,
		s6QueryBuilder: s6QueryBuilder{
			s6Monitor: t4p7s6monitor,
			quote:     t4p7s6monitor.i9Dialect.f8GetQuoter(),
		},
	}
}

func (p7this *S6DeleteBuilder[T]) F8SetEntity(p7Entity *T) *S6DeleteBuilder[T] {
	if nil == p7Entity {
		return p7this
	}
	p7this.p7Entity = p7Entity
	return p7this
}

func (p7this *S6DeleteBuilder[T]) F8Where(s5condition ...S6WhereCondition) *S6DeleteBuilder[T] {
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

func (p7this *S6DeleteBuilder[T]) F8BuildQuery() (*S6Query, error) {
	var err error = nil

	if nil == p7this.p7Entity {
		p7this.p7Entity = new(T)
	}
	if 0 >= len(p7this.s5where) {
		return nil, internal.ErrDeleteWithoutWhere
	}

	p7this.s6QueryBuilder.p7s6Model, err = p7this.s6Monitor.i9Registry.F8Get(p7this.p7Entity)
	if nil != err {
		return nil, err
	}

	p7this.s6QueryBuilder.sqlString.WriteString("DELETE FROM ")
	p7this.s6QueryBuilder.f8WrapWithQuote(p7this.s6QueryBuilder.p7s6Model.TableName)

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

func (p7this *S6DeleteBuilder[T]) F8EXEC(ctx context.Context) (sql.Result, error) {
	p7s6Context := &S6QueryContext{
		QueryType: "DELETE",
		i9Builder: p7this,
		p7s6Model: p7this.s6QueryBuilder.p7s6Model,
		p7s6Query: nil,
	}
	p7s6Result := f8DoEXEC(ctx, p7this.i9Session, &p7this.s6Monitor, p7s6Context)
	return p7s6Result.I9SQLResult, p7s6Result.I9Err
}
