package insert

import (
	"demo-golang/orm"
	"demo-golang/orm/metadata"
)

// MySQL 8 官方文档中，关于 INSERT 语句的部分定义
// https://dev.mysql.com/doc/refman/8.0/en/insert.html
// 13.2.6 INSERT Statement
// INSERT [LOW_PRIORITY | DELAYED | HIGH_PRIORITY] [IGNORE]
//     [INTO] tbl_name
//     [PARTITION (partition_name [, partition_name] ...)]
//     [(col_name [, col_name] ...)]
//     { {VALUES | VALUE} (value_list) [, (value_list)] ... }
//     [AS row_alias[(col_alias [, col_alias] ...)]]
//     [ON DUPLICATE KEY UPDATE assignment_list]

// 这里只处理最常用的部分
// INSERT
//     [INTO] tbl_name
//     [(col_name [, col_name] ...)]
//     { {VALUES | VALUE} (value_list) [, (value_list)] ... }
//     [ON DUPLICATE KEY UPDATE assignment_list]

type S6Insert[T any] struct {
	p7s6OrmDB *S6OrmDB
	// p7s6OrmModel orm 映射模型
	p7s6OrmModel *metadata.S6OrmModel

	s6builder

	// tableName 表名
	tableName string
	// s5column 插入的字段
	s5column []string

	s5value []*T
}

func F8NewS6Insert[T any](p7s6OrmDB *S6OrmDB) *S6Insert[T] {
	return &S6Insert[T]{
		p7s6OrmDB: p7s6OrmDB,
	}
}

func (p7this *S6Insert[T]) Value(s5value ...*T) *S6Insert[T] {
	if 0 >= len(s5value) {
		return p7this
	}

	if nil == p7this.s5value {
		p7this.s5value = s5value
		return p7this
	}
	p7this.s5value = append(p7this.s5value, s5value...)
	return p7this
}

func (p7this *S6Insert[T]) Column(s5column ...string) *S6Insert[T] {
	if 0 >= len(s5column) {
		return p7this
	}

	if nil == p7this.s5column {
		p7this.s5column = s5column
		return p7this
	}
	p7this.s5column = append(p7this.s5column, s5column...)
	return p7this
}

func (p7this *S6Insert[T]) BuildQuery() (*orm.Query, error) {

	return nil, nil
}
