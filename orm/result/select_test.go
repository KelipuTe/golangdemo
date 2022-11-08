package result

import (
	"context"
	orm_select "demo-golang/orm/selectkn"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type S6TestModel struct {
	Id   int
	Name string
	Age  int8
	Sex  int8
}

func TestS6SelectorF8Get(p7s6t *testing.T) {
	// 构造 mock 数据库连接
	p7s6dbMock, sqlMock, err := sqlmock.New()
	if nil != err {
		p7s6t.Fatal(err)
	}
	defer func() {
		_ = p7s6dbMock.Close()
	}()

	p7s6OrmDB := F8NewS6OrmDB(p7s6dbMock)

	s5case := []struct {
		name      string
		sqlString string
		rowsMock  *sqlmock.Rows
		errMock   error
		valueWant *S6TestModel
		errWant   error
	}{
		{
			name:      "normal_sql",
			sqlString: "SELECT .*",
			rowsMock: func() *sqlmock.Rows {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "sex"})
				rows.AddRow([]byte("11"), []byte("aa"), []byte("22"), []byte("1"))
				return rows
			}(),
			valueWant: &S6TestModel{
				Id:   11,
				Name: "aa",
				Age:  22,
				Sex:  1,
			},
		},
	}
	// 把预设的查询结果装进 mock
	for _, t4case := range s5case {
		t4p7eq := sqlMock.ExpectQuery(t4case.sqlString)
		if nil != t4case.errMock {
			t4p7eq.WillReturnError(t4case.errMock)
		} else {
			t4p7eq.WillReturnRows(t4case.rowsMock)
		}
	}

	for _, t4case := range s5case {
		p7s6t.Run(t4case.name, func(p7s6t *testing.T) {
			s6query := orm_select.Query{SQLString: t4case.sqlString, S5parameter: []any{}}
			res, err := F8NewS6OrmSelect[S6TestModel](p7s6OrmDB, s6query).F4Get(context.Background())
			assert.Equal(p7s6t, t4case.errWant, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t, t4case.valueWant, res)
		})
	}
}
