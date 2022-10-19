package orm_select

// Predicate 代表查询语句里的 where、having 子句的查询条件
// Predicate 可以嵌套，组成复杂的查询条件
type Predicate struct {
	// left 操作符左边的查询条件
	left Expression
	// op 中间操作符
	op operator
	// right 操作符右边的查询条件
	right Expression
}

func (this Predicate) doExpression() {}

// And 与，左查询条件 `与` 右查询条件 => (`Id` = 11) AND (Name = 'aa')
func (this Predicate) And(p Predicate) Predicate {
	return Predicate{
		left:  this,
		op:    opAND,
		right: p,
	}
}

// Or 或，左查询条件 `或` 右查询条件 => (`Id` = 11) OR (Name = 'aa')
func (this Predicate) Or(p Predicate) Predicate {
	return Predicate{
		left:  this,
		op:    opOR,
		right: p,
	}
}

// Not 非，`非` 右查询条件 => NOT (`id` = 11)
// 注意 Not 条件只有操作符右边的查询条件
func Not(p Predicate) Predicate {
	return Predicate{
		left:  nil,
		op:    opNOT,
		right: p,
	}
}
