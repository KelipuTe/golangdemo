package orm_select

import "strings"

type OrmSelect struct {
	// s5column 查询的字段
	s5column []string
	// tableName 表名
	tableName string
	// s5where where 语句
	s5where []Predicate
	// s5groupBy group by 语句
	s5groupBy []Column
	// s5having group by 的 having 语句
	s5having []Predicate
	// s5orderBy order by 语句
	s5orderBy []OrderBy
	limit     int
	offset    int
	// 构造出来的 sql
	sqlString strings.Builder
	// sql 中占位符对应的数据
	s5parameter []any
}

// Where 添加 where 子句
func (p7this *OrmSelect) Where(s5w ...Predicate) *OrmSelect {
	if 0 >= len(s5w) {
		return p7this
	}

	if nil == p7this.s5where {
		p7this.s5where = s5w
		return p7this
	}
	p7this.s5where = append(p7this.s5where, s5w...)
	return p7this
}

// GroupBy 添加 group by 子句
func (p7this *OrmSelect) GroupBy(s5c ...Column) *OrmSelect {
	if 0 >= len(s5c) {
		return p7this
	}

	if nil == p7this.s5groupBy {
		p7this.s5groupBy = s5c
		return p7this
	}
	p7this.s5groupBy = append(p7this.s5groupBy, s5c...)
	return p7this
}

// Having 添加 Having 子句
func (p7this *OrmSelect) Having(s5h ...Predicate) *OrmSelect {
	if 0 >= len(s5h) {
		return p7this
	}

	if nil == p7this.s5having {
		p7this.s5having = s5h
		return p7this
	}
	p7this.s5having = append(p7this.s5having, s5h...)
	return p7this
}

func (p7this *OrmSelect) OrderBy(s5ob ...OrderBy) *OrmSelect {
	if 0 >= len(s5ob) {
		return p7this
	}

	if nil == p7this.s5orderBy {
		p7this.s5orderBy = s5ob
		return p7this
	}
	p7this.s5orderBy = append(p7this.s5orderBy, s5ob...)
	return p7this
}

func (p7this *OrmSelect) Limit(l int) *OrmSelect {
	p7this.limit = l
	return p7this
}

func (p7this *OrmSelect) Offset(o int) *OrmSelect {
	p7this.offset = o
	return p7this
}

// addParameter 添加占位符对应的参数
func (p7this *OrmSelect) addParameter(s5p ...any) {
	if nil == p7this.s5parameter {
		p7this.s5parameter = make([]any, 0, 2)
	}
	p7this.s5parameter = append(p7this.s5parameter, s5p...)
}

func (p7this *OrmSelect) BuildQuery() (*Query, error) {
	var err error

	p7this.sqlString.WriteString("SELECT ")

	// 处理查询的列
	p7this.sqlString.WriteString("*")

	p7this.sqlString.WriteString(" FROM ")

	// 处理表名
	p7this.sqlString.WriteByte('`')
	p7this.sqlString.WriteString(p7this.tableName)
	p7this.sqlString.WriteByte('`')

	// 处理 where
	if 0 < len(p7this.s5where) {
		p7this.sqlString.WriteString(" WHERE ")
		err = p7this.buildPredicate(p7this.s5where)
		if nil != err {
			return nil, err
		}
	}

	// 处理 group by
	if 0 < len(p7this.s5groupBy) {
		p7this.sqlString.WriteString(" GROUP BY ")
		for i, t4c := range p7this.s5groupBy {
			if i > 0 {
				p7this.sqlString.WriteByte(',')
			}
			err = p7this.buildColumn(t4c)
			if nil != err {
				return nil, err
			}
		}

		// 在有 group by 的情况下，才处理 having
		if 0 < len(p7this.s5having) {
			p7this.sqlString.WriteString(" HAVING ")
			err = p7this.buildPredicate(p7this.s5having)
			if nil != err {
				return nil, err
			}
		}
	}

	// 处理 order by
	if 0 < len(p7this.s5orderBy) {
		p7this.sqlString.WriteString(" ORDER BY ")
		for i, t4ob := range p7this.s5orderBy {
			if i > 0 {
				p7this.sqlString.WriteByte(',')
			}
			err = p7this.buildColumn(t4ob.column)
			if nil != err {
				return nil, err
			}
			p7this.sqlString.WriteByte(' ')
			p7this.sqlString.WriteString(t4ob.order)
		}
	}

	// 处理 limit offset
	if p7this.limit > 0 {
		p7this.sqlString.WriteString(" LIMIT ?")
		p7this.addParameter(p7this.limit)
	}
	if p7this.offset > 0 {
		p7this.sqlString.WriteString(" OFFSET ?")
		p7this.addParameter(p7this.offset)
	}

	p7this.sqlString.WriteString(";")

	return &Query{
		SQLString:   p7this.sqlString.String(),
		S5parameter: p7this.s5parameter,
	}, nil
}

func (p7this *OrmSelect) buildColumn(c Column) error {
	p7this.sqlString.WriteByte('`')
	p7this.sqlString.WriteString(c.name)
	p7this.sqlString.WriteByte('`')
	return nil
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
		// 处理语句
		t4predicate := e.(Predicate)
		// 递归处理左边的部分
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
		// 递归处理右边的部分
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
		// 处理列名
		t4c := e.(Column)
		err = p7this.buildColumn(t4c)
		if nil != err {
			return err
		}
	case parameter:
		// 处理占位符对应的参数
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
