package orm

import (
	"context"
	"database/sql"
)

// S6Tx 框架的事务对象：封装真正的事务对象
type S6Tx struct {
	// p7s6SqlTx 真正的事务对象
	p7s6SqlTx *sql.Tx
	// p7s6DB 框架的数据库对象
	p7s6DB *S6DB
}

func (p7this *S6Tx) f8GetS6Monitor() s6Monitor {
	return p7this.p7s6DB.s6Monitor
}

func (p7this *S6Tx) f8DoQueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return p7this.p7s6SqlTx.QueryContext(ctx, query, args...)
}

func (p7this *S6Tx) f8DoEXECContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return p7this.p7s6SqlTx.ExecContext(ctx, query, args...)
}

func (p7this *S6Tx) F8Commit() error {
	return p7this.p7s6SqlTx.Commit()
}

func (p7this *S6Tx) F8Rollback() error {
	return p7this.p7s6SqlTx.Rollback()
}

func (p7this *S6Tx) F8RollbackIfNotCommit() error {
	err := p7this.p7s6SqlTx.Rollback()
	if err != sql.ErrTxDone {
		return err
	}
	return nil
}
