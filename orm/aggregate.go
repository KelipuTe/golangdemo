package orm

// #### type ####

// S6Aggregate 对应聚合函数
// 即 SELECT Statement 里的 aggregate
type S6Aggregate struct {
	// name 聚合函数的函数名
	name string
	// s6Column 聚合函数操作的列
	s6Column S6Column
	// alias 别名
	alias string
}

// #### func ####

// F8Count 对列求和
func F8Count(name string) S6Aggregate {
	return S6Aggregate{
		name:     "COUNT",
		s6Column: S6Column{fieldName: name},
	}
}

// F8Avg 对列求平均
func F8Avg(name string) S6Aggregate {
	return S6Aggregate{
		name:     "AVG",
		s6Column: S6Column{fieldName: name},
	}
}

// #### type func ####

// f8BuildAggregate 构造聚合函数 SQL
func (this S6Aggregate) f8BuildAggregate(p7s6Builder *s6QueryBuilder, isUseAlias bool) error {
	p7s6Builder.sqlString.WriteString(this.name)
	p7s6Builder.sqlString.WriteByte('(')
	err := this.s6Column.f8BuildColumn(p7s6Builder, false)
	if nil != err {
		return err
	}
	p7s6Builder.sqlString.WriteByte(')')
	if isUseAlias && "" != this.alias {
		p7s6Builder.sqlString.WriteString(" AS ")
		p7s6Builder.f8WrapWithQuote(this.alias)
	}
	return nil
}

func (this S6Aggregate) f8BuildSelectExpr(p7s6Builder *s6QueryBuilder) error {
	return this.f8BuildAggregate(p7s6Builder, true)
}

func (this S6Aggregate) f8GetFieldName() string {
	return this.s6Column.fieldName
}

func (this S6Aggregate) f8GetAlias() string {
	return this.alias
}

func (this S6Aggregate) f8BuildExpression(p7s6Builder *s6QueryBuilder) error {
	return this.f8BuildAggregate(p7s6Builder, false)
}

// Equal 等于，this = input
func (this S6Aggregate) Equal(input any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorEqual,
		i9RightExpr: f8NewI9Expression(input),
	}
}

// GreaterThan 大于，this > input
func (this S6Aggregate) GreaterThan(input any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorGreaterThan,
		i9RightExpr: f8NewI9Expression(input),
	}
}

// LessThan 小于，this < input
func (this S6Aggregate) LessThan(input any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorLessThan,
		i9RightExpr: f8NewI9Expression(input),
	}
}

// F8As 设置别名
func (this S6Aggregate) F8As(alias string) S6Aggregate {
	return S6Aggregate{
		name:     this.name,
		s6Column: this.s6Column,
		alias:    alias,
	}
}
