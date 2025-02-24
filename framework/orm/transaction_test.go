package orm

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestS6TxF8Commit(p7s6t *testing.T) {
	var err error = nil

	// 构造 mock 数据库连接
	p7s6MockDB, sqlMock, err := sqlmock.New()
	if nil != err {
		p7s6t.Fatal(err)
	}
	defer func() {
		_ = p7s6MockDB.Close()
	}()
	p7s6DB := F8NewS6DB(p7s6MockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectCommit()

	p7s6Tx, err := p7s6DB.F8BeginTx(context.Background(), &sql.TxOptions{})
	assert.Nil(p7s6t, err)

	err = p7s6Tx.F8Commit()
	assert.Nil(p7s6t, err)
}

func TestS6TxF8RollBack(p7s6t *testing.T) {
	var err error = nil

	// 构造 mock 数据库连接
	p7s6MockDB, sqlMock, err := sqlmock.New()
	if nil != err {
		p7s6t.Fatal(err)
	}
	defer func() {
		_ = p7s6MockDB.Close()
	}()
	p7s6DB := F8NewS6DB(p7s6MockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectRollback()

	p7s6Tx, err := p7s6DB.F8BeginTx(context.Background(), &sql.TxOptions{})
	assert.Nil(p7s6t, err)

	err = p7s6Tx.F8Rollback()
	assert.Nil(p7s6t, err)
}
