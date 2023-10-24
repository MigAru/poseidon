package sqlite

import "github.com/huandu/go-sqlbuilder"

func (db *DB) MarkDeleteRepository(id string) error {
	builder := sqlbuilder.SQLite.NewInsertBuilder()

	builder.InsertInto("repository_delete").Cols("repository_id")
	builder.Values(id)

	sqlRaw, args := builder.Build()
	if _, err := db.conn.Exec(sqlRaw, args...); err != nil {
		return err
	}
	return nil
}
