package result

import (
	"database/sql"
	orm_metadata "demo-golang/orm/metadata"
)

// I9Result 接口抽象：用数据库返回的查询结果构造结构体
type I9Result interface {
	// F8SetColumn 方法抽象：将数据库返回的查询结果放到结构体对应的字段上去
	F8SetColumn(rows *sql.Rows) error
}

// F8NewI9Result 方法抽象：创建一个 I9Result 接口的实例
type F8NewI9Result func(value interface{}, p7s5OrmModel *orm_metadata.S6OrmModel) I9Result
