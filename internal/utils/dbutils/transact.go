package dbutils

import (
	"context"
	"database/sql"
)

type TxFunc func(tx *sql.Tx) error

func WithTransaction(
	ctx context.Context,
	db *sql.DB,
	fn TxFunc,
) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = fn(tx)

	return err
}
