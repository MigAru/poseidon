package sqlite

import (
	"github.com/huandu/go-sqlbuilder"
)

func (db *DB) CreateRepository(reference, tag, digest string) error {
	builder := sqlbuilder.SQLite.NewInsertBuilder()
	builder.InsertInto("repository").Cols("reference", "tag", "digest")
	builder.Values(reference, tag, digest)

	sqlRaw, args := builder.Build()

	if _, err := db.conn.Exec(sqlRaw, args...); err != nil {
		return err
	}
	return nil
}
