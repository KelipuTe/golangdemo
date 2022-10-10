package orm_select

// Predicate 查询条件
type Predicate struct {
	left  Expression
	op    operator
	right Expression
}

func (this Predicate) doExpression() {
}

func (this Predicate) And(p Predicate) Predicate {
	return Predicate{
		left:  this,
		op:    opAND,
		right: p,
	}
}

func (this Predicate) Or(p Predicate) Predicate {
	return Predicate{
		left:  this,
		op:    opOR,
		right: p,
	}
}

func Not(p Predicate) Predicate {
	return Predicate{
		op:    opNOT,
		right: p,
	}
}
