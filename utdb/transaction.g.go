package utdb

import (
	"context"
	"database/sql"
)

// BeginTx 开启事务
func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	tx, err := db.DB.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &Tx{tx}, nil
}

// Tx 封装事务
type Tx struct {
	*sql.Tx
}

// Commit 提交事务
func (tx *Tx) Commit() error {
	return tx.Tx.Commit()
}

// Rollback 回滚事务
func (tx *Tx) Rollback() error {
	return tx.Tx.Rollback()
}
