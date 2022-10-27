package orm_result

import (
	"database/sql"
	orm_metadata "demo-golang/orm/orm-metadata"
)

// S6OrmDB 数据库对象
type S6OrmDB struct {
	// 真正的数据库对象
	p7s6SqlDB *sql.DB
	// 元数据
	I9Registry orm_metadata.I9Registry
	// 构造器
	f8NewI9Result F8NewI9Result
}

func F8NewS6OrmDB(p7s6db *sql.DB) *S6OrmDB {
	return &S6OrmDB{
		p7s6SqlDB:     p7s6db,
		I9Registry:    orm_metadata.NewI9Registry(),
		f8NewI9Result: F8NewResultUseReflect,
	}
}
