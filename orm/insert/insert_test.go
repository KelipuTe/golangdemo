package insert

import (
	"demo-golang/orm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestS6OrmInsertF8Build(p7t *testing.T) {
	p7S6OrmDB := F8NewS6OrmDB(nil)
	s5case := []struct {
		name      string
		i9qb      orm.QueryBuilder
		wantQuery *orm.Query
		wantErr   error
	}{
		{
			name: "insert",
			i9qb: F8NewS6Insert[orm.TestModel](p7S6OrmDB),
		},
	}

	for _, t4case := range s5case {
		p7t.Run(t4case.name, func(t *testing.T) {
			query, err := t4case.i9qb.BuildQuery()
			assert.Equal(t, t4case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4case.wantQuery, query)
		})
	}
}
