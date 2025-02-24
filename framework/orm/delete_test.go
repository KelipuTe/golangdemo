package orm

import (
	"demo-golang/orm/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteBuild(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name      string
		i9Builder I9QueryBuilder
		wantQuery *S6Query
		wantErr   error
	}{
		{
			name:      "delete_without_where",
			i9Builder: F8NewS6DeleteBuilder[S6APPUserModel](p7s6DB),
			wantQuery: nil,
			wantErr:   internal.ErrDeleteWithoutWhere,
		},
		{
			name: "delete_one",
			i9Builder: F8NewS6DeleteBuilder[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "DELETE FROM `app_user` WHERE `id` = ?;",
				S5Value:   []any{11},
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
