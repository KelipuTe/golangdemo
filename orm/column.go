package orm

import (
	"demo-golang/orm/internal"
	"fmt"
)

// #### type ####

// S6Column 对应列
// 即各个 Statement 里的 col_name，语句中表示[表、JOIN、子查询]中列的部分
type S6Column struct {
	// i9From 列对应的[表、JOIN、子查询]
	i9From i9TableReference
	// fieldName 结构体属性名
	fieldName string
	// alias 数据库列名的别名
	alias string
}

// #### func ####

func F8NewS6Column(name string) S6Column {
	return S6Column{
		i9From:    nil,
		fieldName: name,
		alias:     "",
	}
}

// #### type func ####

// f8BuildColumn 构造列 SQL
// p7s6Builder 查询构造器
// isUseAlias 用不用别名
func (this S6Column) f8BuildColumn(p7s6Builder *s6QueryBuilder, isUseAlias bool) error {
	// 初始化列名为空，默认找不到列
	var columnName string = ""
	var err error = internal.F8NewErrUnknownField(this.fieldName)

	// 处理表
	if nil != this.i9From {
		columnName, err = this.i9From.f8CheckColumn(p7s6Builder, this)
		// 处理表的别名
		alies := this.i9From.f8GetTableReferenceAlies()
		if "" != alies {
			p7s6Builder.f8WrapWithQuote(alies)
			p7s6Builder.sqlString.WriteByte('.')
		}
	}
	// 上面的逻辑没找到属性，就走默认逻辑，再校验一次
	if nil != err {
		// 校验属性存不存在，存在转换成数据库列名
		p7s6ModelField, ok := p7s6Builder.p7s6Model.M3FieldToColumn[this.fieldName]
		if !ok {
			return internal.F8NewErrUnknownField(this.fieldName)
		}
		columnName = p7s6ModelField.ColumnName
	}
	p7s6Builder.f8WrapWithQuote(columnName)
	// 处理列的别名
	if isUseAlias && "" != this.alias {
		p7s6Builder.sqlString.WriteString(" AS ")
		p7s6Builder.f8WrapWithQuote(this.alias)
	}
	return nil
}

func (this S6Column) f8BuildSelectExpr(p7s6qb *s6QueryBuilder) error {
	return this.f8BuildColumn(p7s6qb, true)
}

func (this S6Column) f8GetFieldName() string {
	return this.fieldName
}

func (this S6Column) f8GetAlias() string {
	return this.alias
}

func (this S6Column) f8BuildExpression(p7s6Builder *s6QueryBuilder) error {
	return this.f8BuildColumn(p7s6Builder, false)
}

// f8BuildAssignment 赋值语句，对应，列 = 列
func (this S6Column) f8BuildAssignment(*s6QueryBuilder) error { return nil }

// ToAssignment 赋值语句，对应，列 = 表达式
func (this S6Column) ToAssignment(input any) S6Assignment {
	i9Expr, ok := input.(i9Expression)
	if !ok {
		i9Expr = S6Value{Value: input}
	}
	return S6Assignment{
		s6Column: this,
		i9Expr:   i9Expr,
	}
}

func (this S6Column) F8Equal(p any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorEqual,
		i9RightExpr: f8NewI9Expression(p),
	}
}

func (this S6Column) F8GreaterThan(p any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorGreaterThan,
		i9RightExpr: f8NewI9Expression(p),
	}
}

func (this S6Column) F8LessThan(p any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorLessThan,
		i9RightExpr: f8NewI9Expression(p),
	}
}

func (this S6Column) F8Like(p any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorLike,
		i9RightExpr: F8NewS6PartRaw(fmt.Sprintf("'%s'", p)),
	}
}

// F8As 设置别名
func (this S6Column) F8As(alias string) S6Column {
	return S6Column{
		i9From:    this.i9From,
		fieldName: this.fieldName,
		alias:     alias,
	}
}

// F8InQuery WHERE IN 查询，列 in (...)
func (this S6Column) F8InQuery(sub S6SubQuery) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorIn,
		i9RightExpr: sub,
	}
}

// F8Add 赋值语句，对应，列 = 列 + n
func (this S6Column) F8Add(num any) S6MathExpression {
	return S6MathExpression{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorAdd,
		i9RightExpr: f8NewI9Expression(num),
	}
}
