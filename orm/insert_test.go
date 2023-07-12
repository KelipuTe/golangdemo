package orm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertBuildMySQL(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name      string
		i9Builder I9QueryBuilder
		wantQuery *S6Query
		wantErr   error
	}{
		{
			name: "insert_one",
			i9Builder: F8NewS6InsertBuilder[S6APPUserModel](p7s6DB).
				F8SetEntity(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}),
			wantQuery: &S6Query{
				SQLString: "INSERT INTO `app_user`(`id`,`name`,`age`,`sex`) VALUES(?,?,?,?);",
				S5Value:   []any{11, "aa", int8(22), int8(1)},
			},
			wantErr: nil,
		},
		{
			name: "insert_two",
			i9Builder: F8NewS6InsertBuilder[S6APPUserModel](p7s6DB).
				F8SetEntity(
					&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1},
					&S6APPUserModel{Id: 22, Name: "bb", Age: 33, Sex: 2},
				),
			wantQuery: &S6Query{
				SQLString: "INSERT INTO `app_user`(`id`,`name`,`age`,`sex`) VALUES(?,?,?,?),(?,?,?,?);",
				S5Value:   []any{11, "aa", int8(22), int8(1), 22, "bb", int8(33), int8(2)},
			},
			wantErr: nil,
		},
		{
			name: "insert_one_set_column",
			i9Builder: F8NewS6InsertBuilder[S6APPUserModel](p7s6DB).
				F8SetEntity(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
				F8SetField("Id", "Name"),
			wantQuery: &S6Query{
				SQLString: "INSERT INTO `app_user`(`id`,`name`) VALUES(?,?);",
				S5Value:   []any{11, "aa"},
			},
			wantErr: nil,
		},
		{
			name: "insert_one_on_conflict_update_with_column",
			i9Builder: F8NewS6InsertBuilder[S6APPUserModel](p7s6DB).
				F8SetEntity(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
				F8OnConflictBuilder().
				F8SetUpdate(F8NewS6Column("Name"), F8NewS6Column("Age")),
			wantQuery: &S6Query{
				SQLString: "INSERT INTO `app_user`(`id`,`name`,`age`,`sex`) VALUES(?,?,?,?) ON DUPLICATE KEY UPDATE `name`=VALUES(`name`),`age`=VALUES(`age`);",
				S5Value:   []any{11, "aa", int8(22), int8(1)},
			},
			wantErr: nil,
		},
		{
			name: "insert_one_on_conflict_update_with_value",
			i9Builder: F8NewS6InsertBuilder[S6APPUserModel](p7s6DB).
				F8SetEntity(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
				F8OnConflictBuilder().
				F8SetUpdate(F8NewS6Column("Name").ToAssignment("bb")),
			wantQuery: &S6Query{
				SQLString: "INSERT INTO `app_user`(`id`,`name`,`age`,`sex`) VALUES(?,?,?,?) ON DUPLICATE KEY UPDATE `name`=?;",
				S5Value:   []any{11, "aa", int8(22), int8(1), "bb"},
			},
			wantErr: nil,
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			query, err := t4value.i9Builder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, query)
		})
	}
}

func TestInsertBuildSQLite3(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil, F8DBWithDialect(S6SQLite3Dialect))

	s5case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "insert_one_on_conflict_update_with_column",
			queryBuilder: F8NewS6InsertBuilder[S6APPUserModel](p7s6DB).
				F8SetEntity(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
				F8OnConflictBuilder().
				F8SetConflictColumn(F8NewS6Column("Id")).
				F8SetUpdate(F8NewS6Column("Name"), F8NewS6Column("Age")),
			wantQuery: &S6Query{
				SQLString: "INSERT INTO `app_user`(`id`,`name`,`age`,`sex`) VALUES(?,?,?,?) ON CONFLICT (`id`) DO UPDATE SET `name`=excluded.`name`,`age`=excluded.`age`;",
				S5Value:   []any{11, "aa", int8(22), int8(1)},
			},
			wantErr: nil,
		},
		{
			name: "insert_one_on_conflict_update_with_value",
			queryBuilder: F8NewS6InsertBuilder[S6APPUserModel](p7s6DB).
				F8SetEntity(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
				F8OnConflictBuilder().
				F8SetConflictColumn(F8NewS6Column("Id")).
				F8SetUpdate(F8NewS6Column("Name").ToAssignment("bb")),
			wantQuery: &S6Query{
				SQLString: "INSERT INTO `app_user`(`id`,`name`,`age`,`sex`) VALUES(?,?,?,?) ON CONFLICT (`id`) DO UPDATE SET `name`=?;",
				S5Value:   []any{11, "aa", int8(22), int8(1), "bb"},
			},
			wantErr: nil,
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, query)
		})
	}
}
