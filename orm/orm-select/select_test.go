package orm_select

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestModel struct {
	Id   int
	Age  int8
	Name string
	Sex  *sql.NullString
}

func TestOrmSelect_BuildQuery(t *testing.T) {
	testCases := []struct {
		name      string
		i9qb      QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			name: "all",
			i9qb: NewOrmSelect(),
			wantQuery: &Query{
				SQLString: "SELECT * FROM `table_name`;",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p7query, err := tc.i9qb.BuildQuery()
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantQuery, p7query)
		})
	}
}

func TestOrmSelect_Where(t *testing.T) {
	testCases := []struct {
		name      string
		i9qb      QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			name: "where one",
			i9qb: NewOrmSelect().
				Where(ToColumn("Id").EQ(11)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE `Id` = ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "where two",
			i9qb: NewOrmSelect().
				Where(ToColumn("Id").EQ(11)).
				Where(ToColumn("Name").EQ("aa")),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND (`Name` =?);",
				S5parameter: []any{11, "aa"},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p7query, err := tc.i9qb.BuildQuery()
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantQuery, p7query)
		})
	}
}
