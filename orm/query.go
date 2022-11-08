package orm

// S6Query 查询语句和参数
// I9QueryBuilder.F8BuildQuery 的结果
type S6Query struct {
	// SQLStr 带有占位符的 SQL 语句
	SQLStr string
	// S5parameter 占位符对应的参数
	S5parameter []any
}

// I9QueryBuilder 查询构造器
// 采用 Builder 设计模式
type I9QueryBuilder interface {
	// F8BuildQuery 构造 SQL
	F8BuildQuery() (*S6Query, error)
}
