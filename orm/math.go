package orm

// S6MathExpression 数学表达式
type S6MathExpression s6BinaryExpression

func (this S6MathExpression) f8BuildExpression(p7s6Builder *s6QueryBuilder) error {
	expr := s6BinaryExpression(this)
	return expr.f8BuildExpression(p7s6Builder)
}
