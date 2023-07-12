package orm

import "demo-golang/orm/internal"

// S6Join 对应 JOIN，可以嵌套
type S6Join struct {
	// i9LeftReference JOIN 左边的部分
	i9LeftReference i9TableReference
	// operator JOIN、LEFT JOIN、RIGHT JOIN
	operator string
	// i9RightReference JOIN 右边的部分
	i9RightReference i9TableReference
	// s5On JOIN ... ON 后面的
	s5On []S6WhereCondition
	// s5Using USING 后面的
	s5Using []S6Column
}

func (this S6Join) f8GetTableReferenceAlies() string {
	return ""
}

func (this S6Join) f8BuildTableReference(p7s6Builder *s6QueryBuilder) error {
	var err error = nil

	p7s6Builder.sqlString.WriteByte('(')

	// 递归处理左边的
	if nil != this.i9LeftReference {
		err = this.i9LeftReference.f8BuildTableReference(p7s6Builder)
		if nil != err {
			return err
		}
	}
	// 处理中间的操作符
	p7s6Builder.sqlString.WriteString(" ")
	p7s6Builder.sqlString.WriteString(this.operator)
	p7s6Builder.sqlString.WriteString(" ")

	// 递归处理右边的
	if nil != this.i9RightReference {
		err = this.i9RightReference.f8BuildTableReference(p7s6Builder)
		if nil != err {
			return err
		}
	}

	if 0 < len(this.s5On) {
		p7s6Builder.sqlString.WriteString(" ON ")
		err = p7s6Builder.f8BuildWhereCondition(this.s5On)
		if nil != err {
			return err
		}
	}

	if 0 < len(this.s5Using) {
		p7s6Builder.sqlString.WriteString(" USING (")
		for i, t4value := range this.s5Using {
			if i > 0 {
				p7s6Builder.sqlString.WriteByte(',')
			}
			err = t4value.f8BuildColumn(p7s6Builder, false)
			if nil != err {
				return err
			}
		}
		p7s6Builder.sqlString.WriteByte(')')
	}

	p7s6Builder.sqlString.WriteByte(')')
	return nil
}

func (this S6Join) f8CheckColumn(p7s6Builder *s6QueryBuilder, s6Column S6Column) (string, error) {
	var columnName string = ""
	var err error = internal.F8NewErrUnknownField(s6Column.fieldName)

	if nil != this.i9LeftReference {
		columnName, err = this.i9LeftReference.f8CheckColumn(p7s6Builder, s6Column)
		if nil == err {
			return columnName, nil
		}
	}
	if nil != this.i9RightReference {
		columnName, err = this.i9RightReference.f8CheckColumn(p7s6Builder, s6Column)
		if nil == err {
			return columnName, nil
		}
	}
	return columnName, err
}

func (this S6Join) F8Join(i9reference i9TableReference) *JoinBuilder {
	return &JoinBuilder{
		i9LeftReference:  this,
		operator:         "JOIN",
		i9RightReference: i9reference,
	}
}

func (this S6Join) F8LeftJoin(i9reference i9TableReference) *JoinBuilder {
	return &JoinBuilder{
		i9LeftReference:  this,
		operator:         "LEFT JOIN",
		i9RightReference: i9reference,
	}
}

type JoinBuilder struct {
	i9LeftReference  i9TableReference
	operator         string
	i9RightReference i9TableReference
}

// F8On 设置 ON
func (this JoinBuilder) F8On(s5condition ...S6WhereCondition) S6Join {
	return S6Join{
		i9LeftReference:  this.i9LeftReference,
		operator:         this.operator,
		i9RightReference: this.i9RightReference,
		s5On:             s5condition,
	}
}

// F8Using 设置 USING
func (this JoinBuilder) F8Using(s5s6Column ...S6Column) S6Join {
	return S6Join{
		i9LeftReference:  this.i9LeftReference,
		operator:         this.operator,
		i9RightReference: this.i9RightReference,
		s5Using:          s5s6Column,
	}
}
