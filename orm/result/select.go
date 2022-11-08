package result

import (
	"context"
	orm_select "demo-golang/orm/selectkn"
)

// S6OrmSelect Select 构造器
type S6OrmSelect[T any] struct {
	p7s6OrmDB *S6OrmDB
	s6query   orm_select.Query
}

// F4Get 执行查询
func (p7this *S6OrmSelect[T]) F4Get(i9ctx context.Context) (*T, error) {
	// 执行查询
	rows, err := p7this.p7s6OrmDB.p7s6SqlDB.QueryContext(i9ctx, p7this.s6query.SQLString, p7this.s6query.S5parameter...)
	if nil != err {
		return nil, err
	}
	// 处理数据库返回的查询结果
	if !rows.Next() {
		return nil, ErrNoRows
	}
	// new 一个类型 T 的变量
	t4p7t := new(T)
	// 获取类型 T 对应的 orm 映射模型
	t4s6OrmModel, err := p7this.p7s6OrmDB.I9Registry.F8Get(t4p7t)
	if nil != err {
		return nil, err
	}
	// 用数据库返回的查询结果构造结构体
	t4result := p7this.p7s6OrmDB.f8NewI9Result(t4p7t, t4s6OrmModel)
	err = t4result.F8SetColumn(rows)

	return t4p7t, err
}

// F8NewS6OrmSelect 构造 S6OrmSelect
func F8NewS6OrmSelect[T any](p7s6db *S6OrmDB, s6query orm_select.Query) *S6OrmSelect[T] {
	return &S6OrmSelect[T]{
		p7s6OrmDB: p7s6db,
		s6query:   s6query,
	}
}
