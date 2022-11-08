package selectkn

// Expression 对应查询语句里的语句或者语句的一部分
type Expression interface {
	doExpression()
}

// toExpression 把输入转换成查询语句里语句
func toExpression(in any) Expression {
	switch in.(type) {
	case Expression:
		// 如果是语句就什么也不做
		return in.(Expression)
	default:
		// 如果不是语句就转换成占位符对应的参数
		return toParameter(in)
	}
}
