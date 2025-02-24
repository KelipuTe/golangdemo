package orm

import (
	"context"
	"demo-golang/orm/internal"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectBuildQuery(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5s6case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name:         "all",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user`;",
			},
		},
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectOperator(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5s6case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "where_equal",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE `id` = ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_equal_use_alias_not_effect",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8As("user_id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE `id` = ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_equal_use_tag",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModelV2](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user_v2` WHERE `user_id` = ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_greater_than",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8GreaterThan(11)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE `id` > ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_less_than",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8LessThan(11)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE `id` < ?;",
				S5Value:   []any{11},
			},
		},
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectWhere(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5s6case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "where_no",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "where_one",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE `id` = ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_two",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11)).
				F8Where(F8NewS6Column("Name").F8Equal("aa")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE (`id` = ?) AND (`name` = ?);",
				S5Value:   []any{11, "aa"},
			},
		},
		{
			name: "where_one_and_one",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11).F8And(F8NewS6Column("Name").F8Equal("aa"))),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE (`id` = ?) AND (`name` = ?);",
				S5Value:   []any{11, "aa"},
			},
		},
		{
			name: "where_one_or_one",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11).F8Or(F8NewS6Column("Name").F8Equal("aa"))),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE (`id` = ?) OR (`name` = ?);",
				S5Value:   []any{11, "aa"},
			},
		},
		{
			name: "where_not_one",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8Not(F8NewS6Column("Id").F8Equal(11))),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE  NOT (`id` = ?);",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_one_and_(one_and_one)",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11).F8And(
					F8NewS6Column("Name").F8Equal("aa").F8And(F8NewS6Column("Age").F8Equal(22)))),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE (`id` = ?) AND ((`name` = ?) AND (`age` = ?));",
				S5Value:   []any{11, "aa", 22},
			},
		},
		{
			name: "where_one_or_(one_or_one)",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11).F8And(
					F8NewS6Column("Name").F8Equal("aa").F8Or(F8NewS6Column("Age").F8Equal(22)))),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE (`id` = ?) AND ((`name` = ?) OR (`age` = ?));",
				S5Value:   []any{11, "aa", 22},
			},
		},
		{
			name: "where_one_and_(not_one)",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11).F8And(
					F8Not(F8NewS6Column("Name").F8Equal("aa")))),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE (`id` = ?) AND ( NOT (`name` = ?));",
				S5Value:   []any{11, "aa"},
			},
		},
		{
			name: "where_one_like",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Name").F8Like("%11%")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE `name` LIKE '%11%';",
				S5Value:   nil,
			},
		},
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectGroupBy(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5s6case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "group_by_no",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8GroupBy(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "group_by_one",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8GroupBy(F8NewS6Column("Age")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` GROUP BY `age`;",
				S5Value:   nil,
			},
		},
		{
			name: "group_by_one_use_alias_not_effect",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8GroupBy(F8NewS6Column("Age").F8As("user_age")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` GROUP BY `age`;",
				S5Value:   nil,
			},
		},
		{
			name: "group_by_one_use_tag",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModelV2](p7s6DB).
				F8GroupBy(F8NewS6Column("Age")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user_v2` GROUP BY `user_age`;",
				S5Value:   nil,
			},
		},
		{
			name: "group_by_two",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8GroupBy(F8NewS6Column("Age")).
				F8GroupBy(F8NewS6Column("Sex")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` GROUP BY `age`,`sex`;",
				S5Value:   nil,
			},
		},
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectHaving(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5s6case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "having_no",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Having(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "having_no_group_by",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Having(F8NewS6Column("Age").F8GreaterThan(22)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "having_one",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8GroupBy(F8NewS6Column("Age")).
				F8Having(F8NewS6Column("Id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` GROUP BY `age` HAVING `id` = ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "group_by_two_having_one",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8GroupBy(F8NewS6Column("Age")).F8GroupBy(F8NewS6Column("Sex")).
				F8Having(F8NewS6Column("Id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` GROUP BY `age`,`sex` HAVING `id` = ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "group_by_two_having_two",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8GroupBy(F8NewS6Column("Age")).F8GroupBy(F8NewS6Column("Sex")).
				F8Having(F8NewS6Column("Id").F8Equal(11)).
				F8Having(F8NewS6Column("Name").F8Equal("aa")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` GROUP BY `age`,`sex` HAVING (`id` = ?) AND (`name` = ?);",
				S5Value:   []any{11, "aa"},
			},
		},
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectOrderBy(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5s6case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "order_by_no",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8OrderBy(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "order_by_one_asc",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8OrderBy(F8Asc("Name")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` ORDER BY `name` ASC;",
				S5Value:   nil,
			},
		},
		{
			name: "order_by_one_desc",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8OrderBy(F8Desc("Name")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` ORDER BY `name` DESC;",
				S5Value:   nil,
			},
		},
		{
			name: "order_by_one_use_tag",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModelV2](p7s6DB).
				F8OrderBy(F8Asc("Age")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user_v2` ORDER BY `user_age` ASC;",
				S5Value:   nil,
			},
		},
		{
			name: "order_by_two_asc_desc",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8OrderBy(F8Asc("Name")).
				F8OrderBy(F8Desc("Age")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` ORDER BY `name` ASC,`age` DESC;",
				S5Value:   nil,
			},
		},
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectOffsetLimit(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5s6case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "limit",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Limit(11),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` LIMIT ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "offset",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Offset(111),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` OFFSET ?;",
				S5Value:   []any{111},
			},
		},
		{
			name: "limit_offset",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Limit(11).F8Offset(111),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` LIMIT ? OFFSET ?;",
				S5Value:   []any{11, 111},
			},
		},
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectSelect(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5s6case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "select_one_column",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Select(F8NewS6Column("Id")),
			wantQuery: &S6Query{
				SQLString: "SELECT `id` FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "select_one_column_use_alias",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Select(F8NewS6Column("Id").F8As("user_id")),
			wantQuery: &S6Query{
				SQLString: "SELECT `id` AS `user_id` FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "select_two_column",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Select(F8NewS6Column("Id")).F8Select(F8NewS6Column("Name")),
			wantQuery: &S6Query{
				SQLString: "SELECT `id`,`name` FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "select_two_column_use_alias",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Select(F8NewS6Column("Id").F8As("user_id")).
				F8Select(F8NewS6Column("Name").F8As("user_name")),
			wantQuery: &S6Query{
				SQLString: "SELECT `id` AS `user_id`,`name` AS `user_name` FROM `app_user`;",
				S5Value:   nil,
			},
		},
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectAggregate(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5s6case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "select_one_aggregate",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Select(F8Count("Id")),
			wantQuery: &S6Query{
				SQLString: "SELECT COUNT(`id`) FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "select_two_aggregate",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Select(F8Count("Id")).F8Select(F8Avg("Age")),
			wantQuery: &S6Query{
				SQLString: "SELECT COUNT(`id`),AVG(`age`) FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "having_one_aggregate",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8GroupBy(F8NewS6Column("Age")).
				F8Having(F8Count("Id").GreaterThan(5)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` GROUP BY `age` HAVING COUNT(`id`) > ?;",
				S5Value:   []any{5},
			},
		},
		{
			name: "select_one_aggregate_use_alias",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Select(F8Count("Id").F8As("count_id")),
			wantQuery: &S6Query{
				SQLString: "SELECT COUNT(`id`) AS `count_id` FROM `app_user`;",
				S5Value:   nil,
			},
		},
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectPartRaw(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5s6case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "select_part_raw",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Select(F8NewS6PartRaw("DISTINCT(`id`)")),
			wantQuery: &S6Query{
				SQLString: "SELECT DISTINCT(`id`) FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "where_part_raw",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6PartRaw("`id` > ?", 11).F8ToWhereCondition()),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE `id` > ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_part_raw_and_one",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6PartRaw("`id` > ?", 11).F8ToWhereCondition().
					F8And(F8NewS6Column("Name").F8Equal("aa"))),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE (`id` > ?) AND (`name` = ?);",
				S5Value:   []any{11, "aa"},
			},
		},
		{
			name: "having_part_raw",
			queryBuilder: F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
				F8GroupBy(F8NewS6Column("Age")).
				F8Having(F8NewS6PartRaw("COUNT(`id`) > ?", 5).F8ToWhereCondition()),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` GROUP BY `age` HAVING COUNT(`id`) > ?;",
				S5Value:   []any{5},
			},
		},
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectJoin(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5s6case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "a_join_b",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserModel{})
				t2 := F8NewS6Table(&S6APPUserInfoModel{})
				return F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
					F8From(t1.F8Join(t2).F8On(t1.F8Column("Id").F8Equal(t2.F8Column("UserID"))))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM (`app_user` JOIN `app_user_info` ON `id` = `user_id`);",
				S5Value:   nil,
			},
		},
		{
			name: "a_join_b_using",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserOrder{})
				t2 := F8NewS6Table(&S6APPUserInfoModel{})
				return F8NewS6SelectBuilder[S6APPUserOrder](p7s6DB).
					F8From(t1.F8Join(t2).F8Using(F8NewS6Column("UserID")))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM (`app_user_order` JOIN `app_user_info` USING (`user_id`));",
				S5Value:   nil,
			},
		},
		{
			name: "a_use_alias_join_b",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserModel{}).F8As("t1")
				t2 := F8NewS6Table(&S6APPUserInfoModel{})
				return F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
					F8From(t1.F8Join(t2).F8On(t1.F8Column("Id").F8Equal(t2.F8Column("UserID"))))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM (`app_user` AS `t1` JOIN `app_user_info` ON `t1`.`id` = `user_id`);",
				S5Value:   nil,
			},
		},
		{
			name: "a_use_alias_join_b_use_alias",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserModel{}).F8As("t1")
				t2 := F8NewS6Table(&S6APPUserInfoModel{}).F8As("t2")
				return F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
					F8From(t1.F8Join(t2).F8On(t1.F8Column("Id").F8Equal(t2.F8Column("UserID"))))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM (`app_user` AS `t1` JOIN `app_user_info` AS `t2` ON `t1`.`id` = `t2`.`user_id`);",
				S5Value:   nil,
			},
		},
		{
			name: "(a_join_b)_join_c",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserModel{}).F8As("t1")
				t2 := F8NewS6Table(&S6APPUserInfoModel{}).F8As("t2")
				t3 := F8NewS6Table(&S6APPUserOrder{}).F8As("t3")
				return F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
					F8From(t1.F8Join(t2).F8On(t1.F8Column("Id").F8Equal(t2.F8Column("UserID"))).
						F8Join(t3).F8On(t1.F8Column("Id").F8Equal(t3.F8Column("UserID"))))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM ((`app_user` AS `t1` JOIN `app_user_info` AS `t2` ON `t1`.`id` = `t2`.`user_id`) JOIN `app_user_order` AS `t3` ON `t1`.`id` = `t3`.`user_id`);",
				S5Value:   nil,
			},
		},
		{
			name: "a_left_join_b",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserModel{}).F8As("t1")
				t2 := F8NewS6Table(&S6APPUserInfoModel{})
				return F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
					F8From(t1.F8LeftJoin(t2).F8On(t1.F8Column("Id").F8Equal(t2.F8Column("UserID"))))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM (`app_user` AS `t1` LEFT JOIN `app_user_info` ON `t1`.`id` = `user_id`);",
				S5Value:   nil,
			},
		},
		{
			name: "(a_left_join_b)_join_c",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserModel{}).F8As("t1")
				t2 := F8NewS6Table(&S6APPUserInfoModel{}).F8As("t2")
				t3 := F8NewS6Table(&S6APPUserOrder{}).F8As("t3")
				return F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).F8From(
					t1.F8LeftJoin(t2).F8On(t1.F8Column("Id").F8Equal(t2.F8Column("UserID"))).
						F8Join(t3).F8On(t1.F8Column("Id").F8Equal(t3.F8Column("UserID"))))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM ((`app_user` AS `t1` LEFT JOIN `app_user_info` AS `t2` ON `t1`.`id` = `t2`.`user_id`) JOIN `app_user_order` AS `t3` ON `t1`.`id` = `t3`.`user_id`);",
				S5Value:   nil,
			},
		},
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectSubQuery(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5s6case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "table_a_from_sub_b",
			queryBuilder: func() I9QueryBuilder {
				sub := F8NewS6SelectBuilder[S6APPUserInfoModel](p7s6DB).F8AsSubQuery("sub")
				return F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).F8From(sub)
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM (SELECT * FROM `app_user_info`) AS `sub`;",
				S5Value:   nil,
			},
		},
		{
			name: "table_a_in_sub_b",
			queryBuilder: func() I9QueryBuilder {
				sub := F8NewS6SelectBuilder[S6APPUserInfoModel](p7s6DB).
					F8Select(F8NewS6Column("UserID")).F8AsSubQuery("sub")
				return F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).
					F8Where(F8NewS6Column("Id").F8InQuery(sub))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE `id` IN (SELECT `user_id` FROM `app_user_info`);",
				S5Value:   nil,
			},
		},
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectJoinAndSubQuery(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5s6case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "table_a_join_sub_b",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserModel{}).F8As("t1")
				subb := F8NewS6SelectBuilder[S6APPUserInfoModel](p7s6DB).F8AsSubQuery("subb")
				return F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).F8From(
					t1.F8Join(subb).F8On(t1.F8Column("Id").F8Equal(subb.F8Column("UserID"))))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM (`app_user` AS `t1` JOIN (SELECT * FROM `app_user_info`) AS `subb` ON `t1`.`id` = `subb`.`user_id`);",
				S5Value:   nil,
			},
		},
		{
			name: "sub_a_join_sub_b",
			queryBuilder: func() I9QueryBuilder {
				suba := F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).F8AsSubQuery("suba")
				subb := F8NewS6SelectBuilder[S6APPUserInfoModel](p7s6DB).F8AsSubQuery("subb")
				return F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).F8From(
					suba.F8Join(subb).F8On(suba.F8Column("Id").F8Equal(subb.F8Column("UserID"))))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM ((SELECT * FROM `app_user`) AS `suba` JOIN (SELECT * FROM `app_user_info`) AS `subb` ON `suba`.`id` = `subb`.`user_id`);",
				S5Value:   nil,
			},
		},
		{
			name: "table_a_join_sub_b_set_select_expr",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserModel{}).F8As("t1")
				subb := F8NewS6SelectBuilder[S6APPUserInfoModel](p7s6DB).
					F8Select(F8NewS6Column("Info")).F8AsSubQuery("subb")
				return F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).F8From(
					t1.F8Join(subb).F8On(t1.F8Column("Id").F8Equal(subb.F8Column("Info"))))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM (`app_user` AS `t1` JOIN (SELECT `info` FROM `app_user_info`) AS `subb` ON `t1`.`id` = `subb`.`info`);",
				S5Value:   nil,
			},
		},
		{
			name: "table_a_join_sub_b_set_select_expr_wrong_column",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserModel{}).F8As("t1")
				subb := F8NewS6SelectBuilder[S6APPUserInfoModel](p7s6DB).
					F8Select(F8NewS6Column("Info2")).F8AsSubQuery("subb")
				return F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).F8From(
					t1.F8Join(subb).F8On(t1.F8Column("Id").F8Equal(subb.F8Column("Info2"))))
			}(),
			wantQuery: nil,
			wantErr:   internal.F8NewErrUnknownField("Info2"),
		},
		{
			name: "select_expr_(table_a_join_sub_b_set_select_expr)",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserModel{}).F8As("t1")
				subb := F8NewS6SelectBuilder[S6APPUserInfoModel](p7s6DB).
					F8Select(F8NewS6Column("Info")).F8AsSubQuery("subb")
				return F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).F8From(
					t1.F8Join(subb).F8On(t1.F8Column("Id").F8Equal(subb.F8Column("Info")))).
					F8Select(t1.F8Column("Name"), subb.F8Column("Info"))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT `t1`.`name`,`subb`.`info` FROM (`app_user` AS `t1` JOIN (SELECT `info` FROM `app_user_info`) AS `subb` ON `t1`.`id` = `subb`.`info`);",
				S5Value:   nil,
			},
		},
		{
			name: "select_expr_wrong_column_(table_a_join_sub_b_set_select_expr)",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserModel{}).F8As("t1")
				subb := F8NewS6SelectBuilder[S6APPUserInfoModel](p7s6DB).
					F8Select(F8NewS6Column("Info")).F8AsSubQuery("subb")
				return F8NewS6SelectBuilder[S6APPUserModel](p7s6DB).F8From(
					t1.F8Join(subb).F8On(t1.F8Column("Id").F8Equal(subb.F8Column("Info")))).
					F8Select(subb.F8Column("Info2"))
			}(),
			wantQuery: nil,
			wantErr:   internal.F8NewErrUnknownField("Info2"),
		},
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectFirst(p7s6t *testing.T) {
	// 构造 mock 数据库连接
	p7s6MockDB, sqlMock, err := sqlmock.New()
	if nil != err {
		p7s6t.Fatal(err)
	}
	defer func() {
		_ = p7s6MockDB.Close()
	}()
	p7s6DB := F8NewS6DB(p7s6MockDB)

	s5s6case := []struct {
		name      string
		sqlString string
		mockRows  *sqlmock.Rows
		mockErr   error
		wantRes   *S6APPUserModel
		wantErr   error
	}{
		{
			name:      "normal_sql",
			sqlString: "SELECT .*",
			mockRows: func() *sqlmock.Rows {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "sex"})
				rows.AddRow([]byte("11"), []byte("aa"), []byte("22"), []byte("1"))
				return rows
			}(),
			wantRes: &S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1},
		},
	}

	// 把预设的查询结果装进 mock
	for _, t4value := range s5s6case {
		t4p7eq := sqlMock.ExpectQuery(t4value.sqlString)
		if nil != t4value.mockErr {
			t4p7eq.WillReturnError(t4value.mockErr)
		} else {
			t4p7eq.WillReturnRows(t4value.mockRows)
		}
	}

	for _, t4value := range s5s6case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			t4p7s6select := F8NewS6SelectBuilder[S6APPUserModel](p7s6DB)
			res, err2 := t4p7s6select.F8First(context.Background())
			assert.Equal(p7s6t2, t4value.wantErr, err2)
			if err2 != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantRes, res)
		})
	}
}
