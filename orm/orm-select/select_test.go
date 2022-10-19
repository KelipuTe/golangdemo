package orm_select

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestModel struct {
	Id   int
	Name string
	Age  int8
	Sex  int8
}

func TestOrmSelect_BuildQuery(t *testing.T) {
	s5case := []struct {
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
	for _, t4c := range s5case {
		t.Run(t4c.name, func(t *testing.T) {
			p7query, err := t4c.i9qb.BuildQuery()
			assert.Equal(t, t4c.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4c.wantQuery, p7query)
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
				Where(NewField("Id").EQ(11)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE `Id` = ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "where_gt",
			i9qb: NewOrmSelect().
				Where(NewField("Id").GT(11)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE `Id` > ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "where_lt",
			i9qb: NewOrmSelect().
				Where(NewField("Id").LT(11)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE `Id` < ?;",
				S5parameter: []any{11},
			},
		},
	}
	for _, t4c := range testCases {
		t.Run(t4c.name, func(t *testing.T) {
			p7query, err := t4c.i9qb.BuildQuery()
			assert.Equal(t, t4c.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4c.wantQuery, p7query)
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
				Where(NewField("Id").EQ(11)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE `Id` = ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "where_two",
			i9qb: NewOrmSelect().
				Where(NewField("Id").EQ(11)).
				Where(NewField("Name").EQ("aa")),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND (`Name` = ?);",
				S5parameter: []any{11, "aa"},
			},
		},
		{
			name: "where_one_and_one",
			i9qb: NewOrmSelect().
				Where(NewField("Id").EQ(11).And(NewField("Name").EQ("aa"))),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND (`Name` = ?);",
				S5parameter: []any{11, "aa"},
			},
		},
		{
			name: "where_one_or_one",
			i9qb: NewOrmSelect().
				Where(NewField("Id").EQ(11).Or(NewField("Name").EQ("aa"))),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) OR (`Name` = ?);",
				S5parameter: []any{11, "aa"},
			},
		},
		{
			name: "where_not_one",
			i9qb: NewOrmSelect().
				Where(Not(NewField("Id").EQ(11))),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE  NOT (`Id` = ?);",
				S5parameter: []any{11},
			},
		},
		{
			name: "where_one_and_(one_and_one)",
			i9qb: NewOrmSelect().
				Where(NewField("Id").EQ(11).And(NewField("Name").EQ("aa").And(NewField("Age").EQ(22)))),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND ((`Name` = ?) AND (`Age` = ?));",
				S5parameter: []any{11, "aa", 22},
			},
		},
		{
			name: "where_one_or_(one_or_one)",
			i9qb: NewOrmSelect().
				Where(NewField("Id").EQ(11).And(NewField("Name").EQ("aa").Or(NewField("Age").EQ(22)))),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND ((`Name` = ?) OR (`Age` = ?));",
				S5parameter: []any{11, "aa", 22},
			},
		},
		{
			name: "where_one_and_(not_one)",
			i9qb: NewOrmSelect().
				Where(NewField("Id").EQ(11).And(Not(NewField("Name").EQ("aa")))),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND ( NOT (`Name` = ?));",
				S5parameter: []any{11, "aa"},
			},
		},
	}
	for _, t4c := range testCases {
		t.Run(t4c.name, func(t *testing.T) {
			p7query, err := t4c.i9qb.BuildQuery()
			assert.Equal(t, t4c.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4c.wantQuery, p7query)
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
			i9qb: NewOrmSelect().
				GroupBy(NewField("Age")),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age`;",
				S5parameter: nil,
			},
		},
		{
			name: "group_by_two",
			i9qb: NewOrmSelect().
				GroupBy(NewField("Age")).GroupBy(NewField("Sex")),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age`,`Sex`;",
				S5parameter: nil,
			},
		},
	}
	for _, t4c := range testCases {
		t.Run(t4c.name, func(t *testing.T) {
			p7query, err := t4c.i9qb.BuildQuery()
			assert.Equal(t, t4c.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4c.wantQuery, p7query)
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
			i9qb: NewOrmSelect().
				Having(NewField("Age").GT(22)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "having_one",
			i9qb: NewOrmSelect().
				GroupBy(NewField("Age")).
				Having(NewField("Id").EQ(11)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age` HAVING `Id` = ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "group_by_two_having_one",
			i9qb: NewOrmSelect().
				GroupBy(NewField("Age")).GroupBy(NewField("Sex")).
				Having(NewField("Id").EQ(11)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age`,`Sex` HAVING `Id` = ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "group_by_two_having_two",
			i9qb: NewOrmSelect().
				GroupBy(NewField("Age")).GroupBy(NewField("Sex")).
				Having(NewField("Id").EQ(11)).Having(NewField("Name").EQ("aa")),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age`,`Sex` HAVING (`Id` = ?) AND (`Name` = ?);",
				S5parameter: []any{11, "aa"},
			},
		},
	}
	for _, t4c := range testCases {
		t.Run(t4c.name, func(t *testing.T) {
			p7query, err := t4c.i9qb.BuildQuery()
			assert.Equal(t, t4c.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4c.wantQuery, p7query)
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
			name: "order_by_one_asc",
			i9qb: NewOrmSelect().
				OrderBy(Asc("Name")),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` ORDER BY `Name` ASC;",
				S5parameter: nil,
			},
		},
		{
			name: "order_by_one_desc",
			i9qb: NewOrmSelect().
				OrderBy(Desc("Name")),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` ORDER BY `Name` DESC;",
				S5parameter: nil,
			},
		},
		{
			name: "order_by_two_asc_desc",
			i9qb: NewOrmSelect().
				OrderBy(Asc("Name")).OrderBy(Desc("Age")),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` ORDER BY `Name` ASC,`Age` DESC;",
				S5parameter: nil,
			},
		},
	}
	for _, t4c := range testCases {
		t.Run(t4c.name, func(t *testing.T) {
			p7query, err := t4c.i9qb.BuildQuery()
			assert.Equal(t, t4c.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4c.wantQuery, p7query)
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
	for _, t4c := range testCases {
		t.Run(t4c.name, func(t *testing.T) {
			p7query, err := t4c.i9qb.BuildQuery()
			assert.Equal(t, t4c.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4c.wantQuery, p7query)
		})
	}
}

func TestOrmSelect_Select(t *testing.T) {
	testCases := []struct {
		name      string
		i9qb      QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			name: "select_one_column",
			i9qb: NewOrmSelect().
				Select(NewField("Id")),
			wantQuery: &Query{
				SQLString:   "SELECT `Id` FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "select_two_column",
			i9qb: NewOrmSelect().
				Select(NewField("Id")).Select(NewField("Name")),
			wantQuery: &Query{
				SQLString:   "SELECT `Id`,`Name` FROM `table_name`;",
				S5parameter: nil,
			},
		},
	}
	for _, t4c := range testCases {
		t.Run(t4c.name, func(t *testing.T) {
			p7query, err := t4c.i9qb.BuildQuery()
			assert.Equal(t, t4c.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4c.wantQuery, p7query)
		})
	}
}

func TestOrmSelect_Aggregate(t *testing.T) {
	testCases := []struct {
		name      string
		i9qb      QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			name: "select_one_aggregate",
			i9qb: NewOrmSelect().
				Select(Count("Id")),
			wantQuery: &Query{
				SQLString:   "SELECT COUNT(`Id`) FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "select_two_aggregate",
			i9qb: NewOrmSelect().
				Select(Count("Id")).Select(Avg("Age")),
			wantQuery: &Query{
				SQLString:   "SELECT COUNT(`Id`),AVG(`Age`) FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "having_one_aggregate",
			i9qb: NewOrmSelect().
				GroupBy(NewField("Age")).
				Having(Count("Id").GT(5)),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age` HAVING COUNT(`Id`) > ?;",
				S5parameter: []any{5},
			},
		},
	}
	for _, t4c := range testCases {
		t.Run(t4c.name, func(t *testing.T) {
			p7query, err := t4c.i9qb.BuildQuery()
			assert.Equal(t, t4c.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4c.wantQuery, p7query)
		})
	}
}

func TestOrmSelect_Raw(t *testing.T) {
	testCases := []struct {
		name      string
		i9qb      QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			name: "select_raw",
			i9qb: NewOrmSelect().
				Select(NewRaw("DISTINCT(Id)")),
			wantQuery: &Query{
				SQLString:   "SELECT DISTINCT(Id) FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "where_raw",
			i9qb: NewOrmSelect().
				Where(NewRaw("Id > ?", 11).toPredicate()),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE Id > ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "where_raw_and_one",
			i9qb: NewOrmSelect().
				Where(NewRaw("Id > ?", 11).toPredicate().And(NewField("Name").EQ("aa"))),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (Id > ?) AND (`Name` = ?);",
				S5parameter: []any{11, "aa"},
			},
		},
		{
			name: "having_raw",
			i9qb: NewOrmSelect().
				GroupBy(NewField("Age")).
				Having(NewRaw("COUNT(Id) > ?", 5).toPredicate()),
			wantQuery: &Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age` HAVING COUNT(Id) > ?;",
				S5parameter: []any{5},
			},
		},
	}
	for _, t4c := range testCases {
		t.Run(t4c.name, func(t *testing.T) {
			p7query, err := t4c.i9qb.BuildQuery()
			assert.Equal(t, t4c.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4c.wantQuery, p7query)
		})
	}
}
