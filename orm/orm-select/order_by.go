package orm_select

// OrderBy 对应查询语句里的 order by
type OrderBy struct {
	// 列
	column Column
	// 排序规则 ASC，DESC
	order string
}

// Asc 升序
func Asc(n string) OrderBy {
	return OrderBy{
		column: Column{name: n},
		order:  "ASC",
	}
}

// Desc 降序
func Desc(n string) OrderBy {
	return OrderBy{
		column: Column{name: n},
		order:  "DESC",
	}
}
