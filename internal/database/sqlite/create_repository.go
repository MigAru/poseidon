package sqlite

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
)

func (db *DB) CreateRepository(tx *sql.Tx, project, tag, digest string) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	if tx == nil {
		defaultTx, err := db.conn.BeginTx(context.Background(), nil)
		if err != nil {
			return err
		}
		tx = defaultTx
	}

	builder := sqlbuilder.SQLite.NewInsertBuilder()
	builder.InsertInto("repository").Cols("id", "project", "tag", "digest")
	builder.Values(id.String(), project, tag, digest)

	sqlRaw, args := builder.Build()

	if _, err := tx.Exec(sqlRaw, args...); err != nil {
		return err
	}
	return nil
}
