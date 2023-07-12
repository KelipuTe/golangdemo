package orm

// S6OrderBy 对应 SELECT 语句的 ORDER BY
// 即 SELECT Statement 里的 ORDER BY {col_name} [ASC | DESC]
type S6OrderBy struct {
	// S6Column 排序列
	S6Column S6Column
	// OrderString 排序规则：ASC；DESC
	OrderString string
}

func (this S6OrderBy) F8BuildOrderBy(p7s6qb *s6QueryBuilder) error {
	err := this.S6Column.f8BuildColumn(p7s6qb, false)
	if nil != err {
		return err
	}
	p7s6qb.sqlString.WriteByte(' ')
	p7s6qb.sqlString.WriteString(this.OrderString)
	return nil
}

// F8Asc 升序
func F8Asc(name string) S6OrderBy {
	return S6OrderBy{
		S6Column:    S6Column{fieldName: name},
		OrderString: "ASC",
	}
}

// F8Desc 降序
func F8Desc(name string) S6OrderBy {
	return S6OrderBy{
		S6Column:    S6Column{fieldName: name},
		OrderString: "DESC",
	}
}
