package selectkn

// Field 对应查询条件里的列
type Field struct {
	// 列名
	name string
}

func (this Field) doExpression() {}

func (this Field) canSelect() {}

func (this Field) EQ(p any) Predicate {
	return Predicate{
		left:  this,
		op:    opEQ,
		right: toExpression(p),
	}
}

func (this Field) GT(p any) Predicate {
	return Predicate{
		left:  this,
		op:    opGT,
		right: toExpression(p),
	}
}

func (this Field) LT(p any) Predicate {
	return Predicate{
		left:  this,
		op:    opLT,
		right: toExpression(p),
	}
}

func NewField(n string) Field {
	return Field{
		name: n,
	}
}
