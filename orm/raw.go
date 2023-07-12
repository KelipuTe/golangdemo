package orm

import (
	"context"
	"database/sql"
)

// S6RawBuilder 原生 SQL 查询构造器
type S6RawBuilder[T any] struct {
	// 带有占位符的 SQL 语句
	sqlString string
	// SQL 语句中占位符对应的参数
	s5Value []any

	i9Session I9Session
	s6QueryBuilder
}

func (p7this *S6RawBuilder[T]) F8BuildQuery() (*S6Query, error) {
	return &S6Query{
		SQLString: p7this.sqlString,
		S5Value:   p7this.s5Value,
	}, nil
}

// F8First 执行查询获取一条数据，用映射关系
func (p7this *S6RawBuilder[T]) F8First(i9ctx context.Context) (*T, error) {
	p7s6Context := &S6QueryContext{
		QueryType: "SELECT",
		i9Builder: p7this,
		p7s6Model: p7this.s6QueryBuilder.p7s6Model,
		p7s6Query: nil,
	}
	p7s6Result := f8DoFirst[T](i9ctx, p7this.i9Session, &p7this.s6Monitor, p7s6Context)
	if nil != p7s6Result.AnyResult {
		return p7s6Result.AnyResult.(*T), p7s6Result.I9Err
	}
	return nil, p7s6Result.I9Err
}

// F8GetList 执行查询获取多条数据，用映射关系
func (p7this *S6RawBuilder[T]) F8GetList(i9ctx context.Context) ([]*T, error) {
	p7s6Context := &S6QueryContext{
		QueryType: "SELECT",
		i9Builder: p7this,
		p7s6Model: p7this.s6QueryBuilder.p7s6Model,
		p7s6Query: nil,
	}
	p7s6Result := f8DoGetList[T](i9ctx, p7this.i9Session, &p7this.s6Monitor, p7s6Context)
	if nil != p7s6Result.AnyResult {
		return p7s6Result.AnyResult.([]*T), p7s6Result.I9Err
	}
	return nil, p7s6Result.I9Err
}

func (p7this *S6RawBuilder[T]) F8EXEC(ctx context.Context) (sql.Result, error) {
	p7s6Context := &S6QueryContext{
		QueryType: "DELETE",
		i9Builder: p7this,
		p7s6Model: p7this.s6QueryBuilder.p7s6Model,
		p7s6Query: nil,
	}
	p7s6Result := f8DoEXEC(ctx, p7this.i9Session, &p7this.s6Monitor, p7s6Context)
	return p7s6Result.I9SQLResult, p7s6Result.I9Err
}

func F8NewS6Raw[T any](i9Session I9Session, sqlString string, s5Value []any) *S6RawBuilder[T] {
	t4p7s6monitor := i9Session.f8GetS6Monitor()
	return &S6RawBuilder[T]{
		i9Session: i9Session,
		s6QueryBuilder: s6QueryBuilder{
			s6Monitor: t4p7s6monitor,
			quote:     t4p7s6monitor.i9Dialect.f8GetQuoter(),
		},
		sqlString: sqlString,
		s5Value:   s5Value,
	}
}
