package insert

import "strings"

type s6builder struct {
	// sqlString 构造出来的 SQL
	sqlString strings.Builder
	// s5parameter SQL 中占位符对应的数据
	s5parameter []any
}
