package orm_result

import (
	"context"
	orm_select "demo-golang/orm/orm-select"
)

// OrmSelect Select 构造器
type OrmSelect[T any] struct {
	p7s6OrmDB *S6OrmDB
	s6query   orm_select.Query
}

func (p7this *OrmSelect[T]) F4Get(i9ctx context.Context) (*T, error) {
	rows, err := p7this.p7s6OrmDB.p7s6SqlDB.QueryContext(i9ctx, p7this.s6query.SQLString, p7this.s6query.S5parameter...)
	if nil != err {
		return nil, err
	}

	if !rows.Next() {
		return nil, ErrNoRows
	}

	t4type := new(T)
	t4s6OrmModel, err := p7this.p7s6OrmDB.I9Registry.F4Get(t4type)
	if nil != err {
		return nil, err
	}
	t4result := p7this.p7s6OrmDB.f8NewI9Result(t4type, t4s6OrmModel)
	err = t4result.F8SetColumn(rows)

	return t4type, err
}

func F4NewOrmSelect[T any](p7s6db *S6OrmDB, s6query orm_select.Query) *OrmSelect[T] {
	return &OrmSelect[T]{
		p7s6OrmDB: p7s6db,
		s6query:   s6query,
	}
}
