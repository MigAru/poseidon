package sqlite

import (
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
)

func (db *DB) CreateRepository(reference, tag, digest string) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	builder := sqlbuilder.SQLite.NewInsertBuilder()
	builder.InsertInto("repository").Cols("id", "reference", "tag", "digest")
	builder.Values(id.String(), reference, tag, digest)

	sqlRaw, args := builder.Build()

	if _, err := db.conn.Exec(sqlRaw, args...); err != nil {
		return err
	}
	return nil
}
