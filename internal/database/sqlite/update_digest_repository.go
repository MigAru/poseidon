package sqlite

import (
	"context"
	"database/sql"
	"github.com/huandu/go-sqlbuilder"
)

func (db *DB) UpdateDigestRepository(tx *sql.Tx, project, tag, digest string) error {
	if tx == nil {
		defaultTx, err := db.conn.BeginTx(context.Background(), nil)
		if err != nil {
			return err
		}
		tx = defaultTx
	}

	builder := sqlbuilder.SQLite.NewUpdateBuilder()
	builder.Update("repository")
	builder.Set(builder.Assign("digest", digest))
	builder.Where(builder.Equal("project", project), builder.Equal("tag", tag))

	sqlRaw, args := builder.Build()
	if _, err := tx.Exec(sqlRaw, args...); err != nil {
		return err
	}
	return nil
}
