package orm_select

// Column 对应查询条件里的列
type Column struct {
	// 列名
	name string
}

func (this Column) doExpression() {}

func (this Column) canSelect() {}

func ToColumn(n string) Column {
	return Column{
		name: n,
	}
}

func (this Column) EQ(p any) Predicate {
	return Predicate{
		left:  this,
		op:    opEQ,
		right: toExpression(p),
	}
}

func (this Column) GT(p any) Predicate {
	return Predicate{
		left:  this,
		op:    opGT,
		right: toExpression(p),
	}
}

func (this Column) LT(p any) Predicate {
	return Predicate{
		left:  this,
		op:    opLT,
		right: toExpression(p),
	}
}
