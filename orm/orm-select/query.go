package orm_select

type QueryBuilder interface {
	BuildQuery() (*Query, error)
}

type Query struct {
	SQLString   string
	S5parameter []any
}
