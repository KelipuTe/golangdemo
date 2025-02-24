package orm

// #### Part Raw ####

// S6PartRaw 对应一段原生 SQL
type S6PartRaw struct {
	// sqlString 带有占位符的 SQL 语句
	sqlString string
	// s5Value SQL 语句中占位符对应的值
	s5Value []any
}

// f8PartRaw 构造一段原生 SQL 对应的 SQL
func (this S6PartRaw) f8PartRaw(p7s6qb *s6QueryBuilder) error {
	p7s6qb.sqlString.WriteString(this.sqlString)
	if 0 < len(this.s5Value) {
		p7s6qb.f8AddParameter(this.s5Value...)
	}
	return nil
}

func (this S6PartRaw) f8BuildSelectExpr(p7s6qb *s6QueryBuilder) error {
	return this.f8PartRaw(p7s6qb)
}

func (this S6PartRaw) f8GetFieldName() string {
	return ""
}

func (this S6PartRaw) f8GetAlias() string {
	return ""
}

func (this S6PartRaw) f8BuildExpression(p7s6Builder *s6QueryBuilder) error {
	return this.f8PartRaw(p7s6Builder)
}

// F8ToWhereCondition 转换成查询条件
func (this S6PartRaw) F8ToWhereCondition() S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  "",
		i9RightExpr: nil,
	}
}

func F8NewS6PartRaw(sql string, s5Value ...any) S6PartRaw {
	return S6PartRaw{
		sqlString: sql,
		s5Value:   s5Value,
	}
}
