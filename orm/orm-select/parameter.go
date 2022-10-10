package orm_select

// value where、having 子句的列名对应的数据
type parameter struct {
	parameter any
}

func (this parameter) doExpression() {
}

func toParameter(p any) parameter {
	return parameter{
		parameter: p,
	}
}
