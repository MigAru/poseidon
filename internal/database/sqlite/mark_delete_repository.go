package sqlite

import (
	"database/sql"
	"github.com/huandu/go-sqlbuilder"
)

func (db *DB) MarkDeleteRepository(tx *sql.Tx, id string) error {
	builder := sqlbuilder.SQLite.NewInsertBuilder()

	builder.InsertInto("repository_delete").Cols("repository_id")
	builder.Values(id)

	sqlRaw, args := builder.Build()
	if _, err := tx.Exec(sqlRaw, args...); err != nil {
		return err
	}
	return nil
}
