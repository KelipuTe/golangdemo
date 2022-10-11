package orm_select

// Aggregate 聚合函数
type Aggregate struct {
	// 聚合函数名
	funcName string
	// 列
	column Column
}

func (this Aggregate) doExpression() {
}

func (this Aggregate) EQ(p any) Predicate {
	return Predicate{
		left:  this,
		op:    opEQ,
		right: toExpression(p),
	}
}

func (this Aggregate) GT(p any) Predicate {
	return Predicate{
		left:  this,
		op:    opGT,
		right: toExpression(p),
	}
}

func (this Aggregate) LT(p any) Predicate {
	return Predicate{
		left:  this,
		op:    opLT,
		right: toExpression(p),
	}
}

func Count(n string) Aggregate {
	return Aggregate{
		funcName: "COUNT",
		column:   Column{name: n},
	}
}

func Sum(n string) Aggregate {
	return Aggregate{
		funcName: "SUM",
		column:   Column{name: n},
	}
}
