package sqlite

import (
	"context"
	"database/sql"
)

func (db *DB) NewTx(ctx context.Context) (*sql.Tx, error) {
	return db.conn.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
}
