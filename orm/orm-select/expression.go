package orm_select

type Expression interface {
	doExpression()
}

func toExpression(in any) Expression {
	switch in.(type) {
	case Expression:
		return in.(Expression)
	default:
		return toParameter(in)
	}
}
