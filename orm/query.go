package orm

import (
	"demo-golang/orm/metadata"
	"strings"
)

// S6Query 构造出来的查询语句和参数
type S6Query struct {
	// SQLString 带有占位符的 SQL 语句
	SQLString string
	// S5Value SQL 语句中占位符对应的值
	S5Value []any
}

// I9QueryBuilder 接口抽象：查询构造器
// Builder 设计模式
type I9QueryBuilder interface {
	// F8BuildQuery 方法抽象：构造 S6Query
	F8BuildQuery() (*S6Query, error)
}

// s6QueryBuilder 查询构造器
type s6QueryBuilder struct {
	// s6Monitor 控制器。从 S6DB 里面获取。
	s6Monitor
	// p7s6Model 查询对应的源数据
	p7s6Model *metadata.S6Model
	// quote 这个东西 s6Monitor 里面有，拿到这里方便操作
	quote byte
	// sqlString 带有占位符的 SQL 语句
	sqlString strings.Builder
	// s5Value SQL 语句中占位符对应的参数
	s5Value []any
}

// f8WrapWithQuote 两边加引号
func (p7this *s6QueryBuilder) f8WrapWithQuote(name string) {
	p7this.sqlString.WriteByte(p7this.quote)
	p7this.sqlString.WriteString(name)
	p7this.sqlString.WriteByte(p7this.quote)
}

// f8BuildWhereCondition 处理查询条件
func (p7this *s6QueryBuilder) f8BuildWhereCondition(s5p []S6WhereCondition) error {
	t4p := s5p[0]
	for i := 1; i < len(s5p); i++ {
		t4p = t4p.F8And(s5p[i])
	}
	return p7this.f8BuildExpression(t4p)
}

func (p7this *s6QueryBuilder) f8BuildExpression(expr i9Expression) error {
	if nil == expr {
		return nil
	}
	return expr.f8BuildExpression(p7this)
}

// f8AddParameter 添加占位符对应的参数
func (p7this *s6QueryBuilder) f8AddParameter(s5p ...any) {
	if nil == p7this.s5Value {
		p7this.s5Value = make([]any, 0, 8)
	}
	p7this.s5Value = append(p7this.s5Value, s5p...)
}
