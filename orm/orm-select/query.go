package orm_select

// QueryBuilder Builder 设计模式
type QueryBuilder interface {
	// BuildQuery 构造 SQL
	BuildQuery() (*Query, error)
}

// Query QueryBuilder.BuildQuery 的结果
type Query struct {
	// 带有占位符的 SQL 语句
	SQLString string
	// 占位符对应的参数
	S5parameter []any
}
