package orm

import "demo-golang/orm/internal"

// S6SubQuery 对应子查询
type S6SubQuery struct {
	// s5Select SELECT 后面的
	s5Select []i9SelectExpr
	// i9From FROM 后面的
	i9From i9TableReference
	// alias 别名
	alias string
	// 构造子查询语句
	i9Builder I9QueryBuilder
}

// F8BuildSubQuery 构造子查询语句
func (this S6SubQuery) F8BuildSubQuery(p7s6Builder *s6QueryBuilder, isUseAlias bool) error {
	query, err := this.i9Builder.F8BuildQuery()
	if err != nil {
		return err
	}
	p7s6Builder.sqlString.WriteByte('(')
	p7s6Builder.sqlString.WriteString(query.SQLString[:len(query.SQLString)-1])
	if 0 < len(query.S5Value) {
		p7s6Builder.f8AddParameter(query.S5Value...)
	}
	p7s6Builder.sqlString.WriteByte(')')
	if isUseAlias && "" != this.alias {
		p7s6Builder.sqlString.WriteString(" AS ")
		p7s6Builder.f8WrapWithQuote(this.alias)
	}
	return nil
}

func (this S6SubQuery) f8BuildExpression(p7s6Builder *s6QueryBuilder) error {
	return this.F8BuildSubQuery(p7s6Builder, false)
}

func (this S6SubQuery) f8BuildTableReference(p7s6Builder *s6QueryBuilder) error {
	return this.F8BuildSubQuery(p7s6Builder, true)
}

func (this S6SubQuery) f8CheckColumn(p7s6Builder *s6QueryBuilder, s6Column S6Column) (string, error) {
	var columnName string = ""
	var err error = internal.F8NewErrUnknownField(s6Column.fieldName)

	if 0 < len(this.s5Select) {
		// 如果设置了查询表达式，就校验在不在设置的查询表达式里
		for _, t4value := range this.s5Select {
			if t4value.f8GetFieldName() == s6Column.fieldName {
				// 如果命中的是结构体属性，就二次校验一下
				columnName, err = this.i9From.f8CheckColumn(p7s6Builder, s6Column)
				if nil == err {
					return columnName, nil
				}
			}
			if t4value.f8GetAlias() == s6Column.fieldName {
				// 如果命中的是别名，那就直接过
				return s6Column.fieldName, nil
			}
		}
		// 如果遍历完还没找到就要报错了，这里不能继续往下走
		return columnName, err
	}
	// 如果没有设置查询表达式，那默认是全要的
	if nil != this.i9From {
		columnName, err = this.i9From.f8CheckColumn(p7s6Builder, s6Column)
		if nil == err {
			return columnName, nil
		}
	}
	return columnName, err
}

func (this S6SubQuery) f8GetTableReferenceAlies() string {
	return this.alias
}

func (this S6SubQuery) F8Join(i9reference i9TableReference) *JoinBuilder {
	return &JoinBuilder{
		i9LeftReference:  this,
		operator:         "JOIN",
		i9RightReference: i9reference,
	}
}

func (this S6SubQuery) F8LeftJoin(i9reference i9TableReference) *JoinBuilder {
	return &JoinBuilder{
		i9LeftReference:  this,
		operator:         "LEFT JOIN",
		i9RightReference: i9reference,
	}
}

func (this S6SubQuery) F8Column(name string) S6Column {
	return S6Column{
		i9From:    this,
		fieldName: name,
	}
}
