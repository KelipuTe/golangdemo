package orm_select

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestModel struct {
	Id   int
	Name string
	Age  int8
	Sex  string
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

func TestOrmSelect_Operator(t *testing.T) {
	testCases := []struct {
		name      string
		i9qb      QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			name: "where_eq",
			i9qb: NewOrmSelect().
				Where(ToColumn("Id").EQ(11)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE `Id` = ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "where_gt",
			i9qb: NewOrmSelect().
				Where(ToColumn("Id").GT(11)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE `Id` > ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "where_lt",
			i9qb: NewOrmSelect().
				Where(ToColumn("Id").LT(11)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE `Id` < ?;",
				S5parameter: []any{11},
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
			name: "where_no",
			i9qb: NewOrmSelect().Where(),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "where_one",
			i9qb: NewOrmSelect().
				Where(ToColumn("Id").EQ(11)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE `Id` = ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "where_two",
			i9qb: NewOrmSelect().
				Where(ToColumn("Id").EQ(11)).
				Where(ToColumn("Name").EQ("aa")),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND (`Name` = ?);",
				S5parameter: []any{11, "aa"},
			},
		},
		{
			name: "where_one_and_one",
			i9qb: NewOrmSelect().
				Where(ToColumn("Id").EQ(11).And(ToColumn("Name").EQ("aa"))),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND (`Name` = ?);",
				S5parameter: []any{11, "aa"},
			},
		},
		{
			name: "where_one_or_one",
			i9qb: NewOrmSelect().
				Where(ToColumn("Id").EQ(11).Or(ToColumn("Name").EQ("aa"))),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) OR (`Name` = ?);",
				S5parameter: []any{11, "aa"},
			},
		},
		{
			name: "where_not_one",
			i9qb: NewOrmSelect().
				Where(Not(ToColumn("Id").EQ(11))),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE  NOT (`Id` = ?);",
				S5parameter: []any{11},
			},
		},
		{
			name: "where_one_and_(one_and_one)",
			i9qb: NewOrmSelect().
				Where(ToColumn("Id").EQ(11).And(ToColumn("Name").EQ("aa").And(ToColumn("Age").EQ(22)))),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND ((`Name` = ?) AND (`Age` = ?));",
				S5parameter: []any{11, "aa", 22},
			},
		},
		{
			name: "where_one_or_(one_or_one)",
			i9qb: NewOrmSelect().
				Where(ToColumn("Id").EQ(11).And(ToColumn("Name").EQ("aa").Or(ToColumn("Age").EQ(22)))),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND ((`Name` = ?) OR (`Age` = ?));",
				S5parameter: []any{11, "aa", 22},
			},
		},
		{
			name: "where_one_and_(not_one)",
			i9qb: NewOrmSelect().
				Where(ToColumn("Id").EQ(11).And(Not(ToColumn("Name").EQ("aa")))),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND ( NOT (`Name` = ?));",
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

func TestOrmSelect_GroupBy(t *testing.T) {
	testCases := []struct {
		name      string
		i9qb      QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			name: "group_by_no",
			i9qb: NewOrmSelect().GroupBy(),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "group_by_one",
			i9qb: NewOrmSelect().GroupBy(ToColumn("Age")),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age`;",
				S5parameter: nil,
			},
		},
		{
			name: "group_by_two",
			i9qb: NewOrmSelect().GroupBy(ToColumn("Age")).GroupBy(ToColumn("Sex")),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age`,`Sex`;",
				S5parameter: nil,
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

func TestOrmSelect_Having(t *testing.T) {
	testCases := []struct {
		name      string
		i9qb      QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			name: "having_no",
			i9qb: NewOrmSelect().Having(),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "having_no_group_by",
			i9qb: NewOrmSelect().Having(ToColumn("Age").GT(22)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "having_one",
			i9qb: NewOrmSelect().GroupBy(ToColumn("Age")).Having(ToColumn("Id").EQ(11)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age` HAVING `Id` = ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "group_by_two_having_one",
			i9qb: NewOrmSelect().
				GroupBy(ToColumn("Age")).GroupBy(ToColumn("Sex")).
				Having(ToColumn("Id").EQ(11)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age`,`Sex` HAVING `Id` = ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "group_by_two_having_two",
			i9qb: NewOrmSelect().
				GroupBy(ToColumn("Age")).GroupBy(ToColumn("Sex")).
				Having(ToColumn("Id").EQ(11)).Having(ToColumn("Name").EQ("aa")),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age`,`Sex` HAVING (`Id` = ?) AND (`Name` = ?);",
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

func TestOrmSelect_OrderBy(t *testing.T) {
	testCases := []struct {
		name      string
		i9qb      QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			name: "order_by_no",
			i9qb: NewOrmSelect().OrderBy(),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "having_no_group_by",
			i9qb: NewOrmSelect().Having(ToColumn("Age").GT(22)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "having_one",
			i9qb: NewOrmSelect().GroupBy(ToColumn("Age")).Having(ToColumn("Id").EQ(11)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age` HAVING `Id` = ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "group_by_two_having_one",
			i9qb: NewOrmSelect().
				GroupBy(ToColumn("Age")).GroupBy(ToColumn("Sex")).
				Having(ToColumn("Id").EQ(11)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age`,`Sex` HAVING `Id` = ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "group_by_two_having_two",
			i9qb: NewOrmSelect().
				GroupBy(ToColumn("Age")).GroupBy(ToColumn("Sex")).
				Having(ToColumn("Id").EQ(11)).Having(ToColumn("Name").EQ("aa")),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age`,`Sex` HAVING (`Id` = ?) AND (`Name` = ?);",
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

func TestOrmSelect_OffsetLimit(t *testing.T) {
	testCases := []struct {
		name      string
		i9qb      QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			name: "limit",
			i9qb: NewOrmSelect().Limit(11),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` LIMIT ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "offset",
			i9qb: NewOrmSelect().Offset(111),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` OFFSET ?;",
				S5parameter: []any{111},
			},
		},
		{
			name: "limit_offset",
			i9qb: NewOrmSelect().Limit(11).Offset(111),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` LIMIT ? OFFSET ?;",
				S5parameter: []any{11, 111},
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
