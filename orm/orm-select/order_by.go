package orm_select

type OrderBy struct {
	// 列
	column Column
	// 排序规则 ASC，DESC
	order string
}

func Asc(n string) OrderBy {
	return OrderBy{
		column: Column{name: n},
		order:  "ASC",
	}
}

func Desc(n string) OrderBy {
	return OrderBy{
		column: Column{name: n},
		order:  "DESC",
	}
}
