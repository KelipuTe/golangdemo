package orm_select

import "strings"

type OrmSelect struct {
	// 查询的字段
	s5column []string
	// 表名
	tableName string
	// where
	s5where   []Predicate
	s5groupBy []Column
	s5having  []Predicate
	s5orderBy []OrderBy
	offset    int
	limit     int
	// 构造出来的 sql
	sqlString strings.Builder
	// sql 中占位符对应的数据
	s5parameter []any
}

func (p7this *OrmSelect) Where(s5where ...Predicate) *OrmSelect {
	if nil == p7this.s5where {
		p7this.s5where = s5where
		return p7this
	}
	p7this.s5where = append(p7this.s5where, s5where...)
	return p7this
}

func (p7this *OrmSelect) addParameter(s5p ...any) {
	if nil == p7this.s5parameter {
		p7this.s5parameter = make([]any, 0, 2)
	}
	p7this.s5parameter = append(p7this.s5parameter, s5p...)
}

func (p7this *OrmSelect) BuildQuery() (*Query, error) {
	var err error

	p7this.sqlString.WriteString("SELECT ")

	p7this.sqlString.WriteString("*")

	p7this.sqlString.WriteString(" FROM ")

	p7this.sqlString.WriteByte('`')
	p7this.sqlString.WriteString(p7this.tableName)
	p7this.sqlString.WriteByte('`')

	if 0 < len(p7this.s5where) {
		p7this.sqlString.WriteString(" WHERE ")
		err = p7this.buildPredicate(p7this.s5where)
		if nil != err {
			return nil, err
		}
	}

	p7this.sqlString.WriteString(";")

	return &Query{
		SQLString:   p7this.sqlString.String(),
		S5parameter: p7this.s5parameter,
	}, nil
}

func (p7this *OrmSelect) buildPredicate(s5p []Predicate) error {
	t4p := s5p[0]
	for i := 1; i < len(s5p); i++ {
		t4p = t4p.And(s5p[i])
	}
	return p7this.buildExpression(t4p)
}

func (p7this *OrmSelect) buildExpression(e Expression) error {
	var err error

	if nil == e {
		return nil
	}

	switch e.(type) {
	case Predicate:
		t4predicate := e.(Predicate)
		// 处理左边的部分
		_, lIsP := t4predicate.left.(Predicate)
		if lIsP {
			p7this.sqlString.WriteByte('(')
		}
		err = p7this.buildExpression(t4predicate.left)
		if nil != err {
			return err
		}
		if lIsP {
			p7this.sqlString.WriteByte(')')
		}
		// 处理中间的操作符
		p7this.sqlString.WriteByte(' ')
		p7this.sqlString.WriteString(t4predicate.op.String())
		p7this.sqlString.WriteByte(' ')
		// 处理右边的部分
		_, rIsP := t4predicate.right.(Predicate)
		if rIsP {
			p7this.sqlString.WriteByte('(')
		}
		err = p7this.buildExpression(t4predicate.right)
		if nil != err {
			return err
		}
		if rIsP {
			p7this.sqlString.WriteByte(')')
		}
	case Column:
		t4column := e.(Column)
		p7this.sqlString.WriteByte('`')
		p7this.sqlString.WriteString(t4column.name)
		p7this.sqlString.WriteByte('`')
	case parameter:
		t4parameter := e.(parameter)
		p7this.sqlString.WriteByte('?')
		p7this.addParameter(t4parameter.parameter)
	default:
		return NewErrUnsupportedExpressionType(e)
	}
	return nil
}

func NewOrmSelect() *OrmSelect {
	return &OrmSelect{
		tableName: "table_name",
	}
}
