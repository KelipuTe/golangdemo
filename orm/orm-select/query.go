package orm_select

// QueryBuilder 查询构造器
// 采用 Builder 设计模式
type QueryBuilder interface {
	// BuildQuery 构造 SQL
	BuildQuery() (*Query, error)
}

// Query QueryBuilder.BuildQuery 的结果
type Query struct {
	// SQLString 带有占位符的 SQL 语句
	SQLString string
	// S5parameter 占位符对应的参数
	S5parameter []any
}
