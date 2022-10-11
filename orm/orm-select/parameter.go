package orm_select

// parameter 对应查询语句里的占位符对应的参数
type parameter struct {
	parameter any
}

func (this parameter) doExpression() {
}

// toParameter 把输入转换成查询语句里的占位符对应的参数
func toParameter(p any) parameter {
	return parameter{
		parameter: p,
	}
}
