package orm_select

// Column 对应查询条件里的列名
type Column struct {
	name string
}

func (this Column) doExpression() {
}

func ToColumn(name string) Column {
	return Column{
		name: name,
	}
}

func (this Column) EQ(p any) Predicate {
	return Predicate{
		left:  this,
		op:    opEQ,
		right: toExpression(p),
	}
}
