package orm

import (
	"demo-golang/orm/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateBuild(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name      string
		i9Builder I9QueryBuilder
		wantQuery *S6Query
		wantErr   error
	}{
		{
			name: "update_without_column",
			i9Builder: F8NewS6UpdateBuilder[S6APPUserModel](p7s6DB).
				F8SetEntity(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}),
			wantQuery: nil,
			wantErr:   internal.ErrUpdateWithoutColumn,
		},
		{
			name: "update_without_where",
			i9Builder: F8NewS6UpdateBuilder[S6APPUserModel](p7s6DB).
				F8SetEntity(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
				F8SetUpdate(F8NewS6Column("Name")),
			wantQuery: nil,
			wantErr:   internal.ErrUpdateWithoutWhere,
		},
		{
			name: "update_one",
			i9Builder: F8NewS6UpdateBuilder[S6APPUserModel](p7s6DB).
				F8SetEntity(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
				F8SetUpdate(F8NewS6Column("Name")).
				F8Where(F8NewS6Column("Id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "UPDATE `app_user` SET `name`=? WHERE `id` = ?;",
				S5Value:   []any{"aa", 11},
			},
			wantErr: nil,
		},
		{
			name: "update_two",
			i9Builder: F8NewS6UpdateBuilder[S6APPUserModel](p7s6DB).
				F8SetEntity(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
				F8SetUpdate(F8NewS6Column("Name"), F8NewS6Column("Age")).
				F8Where(F8NewS6Column("Id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "UPDATE `app_user` SET `name`=?,`age`=? WHERE `id` = ?;",
				S5Value:   []any{"aa", int8(22), 11},
			},
			wantErr: nil,
		},
		{
			name: "update_one=one+1",
			i9Builder: F8NewS6UpdateBuilder[S6APPUserModel](p7s6DB).
				F8SetEntity(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
				F8SetUpdate(F8NewS6Column("Age").ToAssignment(F8NewS6Column("Age").F8Add(1))).
				F8Where(F8NewS6Column("Id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "UPDATE `app_user` SET `age`=`age` + ? WHERE `id` = ?;",
				S5Value:   []any{1, 11},
			},
			wantErr: nil,
		},
		{
			name: "update_one_part_raw",
			i9Builder: F8NewS6UpdateBuilder[S6APPUserModel](p7s6DB).
				F8SetEntity(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
				F8SetUpdate(F8NewS6Column("Age").ToAssignment(F8NewS6PartRaw("`age`+1"))).
				F8Where(F8NewS6Column("Id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "UPDATE `app_user` SET `age`=`age`+1 WHERE `id` = ?;",
				S5Value:   []any{11},
			},
			wantErr: nil,
		},
		{
			name: "update_one_without_entity",
			i9Builder: F8NewS6UpdateBuilder[S6APPUserModel](p7s6DB).
				F8SetUpdate(F8NewS6Column("Name").ToAssignment("bb")).
				F8Where(F8NewS6Column("Id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "UPDATE `app_user` SET `name`=? WHERE `id` = ?;",
				S5Value:   []any{"bb", 11},
			},
			wantErr: nil,
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			query, err := t4value.i9Builder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if nil != err {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, query)
		})
	}
}
