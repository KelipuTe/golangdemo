package orm

import (
	"database/sql"
	"fmt"
	"testing"
)

type TestModel struct {
	Id        int64
	FirstName string
	Age       int8
	LastName  *sql.NullString
}

// memoryDB 返回一个基于内存的 ORM，它使用的是 sqlite3 内存模式。
func memoryDB(t *testing.T) *DB {
	orm, err := Open("sqlite3", "file:test.db?cache=shared&mode=memory")
	if err != nil {
		t.Fatal(err)
	}
	return orm
}

func memoryDBWithDB(db string, t *testing.T) *DB {
	orm, err := Open("sqlite3", fmt.Sprintf("file:%s.db?cache=shared&mode=memory", db))
	if err != nil {
		t.Fatal(err)
	}
	return orm
}
