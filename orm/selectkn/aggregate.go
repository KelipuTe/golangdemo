package selectkn

// Aggregate 对应查询语句里的聚合函数
type Aggregate struct {
	// funcName 聚合函数名
	funcName string
	// field 列
	field Field
}

func (this Aggregate) doExpression() {}

func (this Aggregate) canSelect() {}

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
		field:    Field{name: n},
	}
}

func Avg(n string) Aggregate {
	return Aggregate{
		funcName: "AVG",
		field:    Field{name: n},
	}
}
