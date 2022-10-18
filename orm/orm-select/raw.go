package orm_select

// Raw 对应原生 sql
type Raw struct {
	// 原生 sql
	raw string
	// sql 中占位符对应的数据
	s5parameter []any
}

func (this Raw) doExpression() {}

func (this Raw) canSelect() {}

func (this Raw) toPredicate() Predicate {
	return Predicate{
		left:  this,
		op:    "",
		right: nil,
	}
}

func ToRaw(raw string, s5p ...any) *Raw {
	return &Raw{
		raw:         raw,
		s5parameter: s5p,
	}
}
