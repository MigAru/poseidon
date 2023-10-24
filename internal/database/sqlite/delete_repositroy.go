package sqlite

import "github.com/huandu/go-sqlbuilder"

func (db *DB) DeleteRepository(id string) error {
	builder := sqlbuilder.SQLite.NewDeleteBuilder()
	builder.DeleteFrom("repository")
	builder.Where(builder.Equal("id", id))

	sqlRaw, args := builder.Build()

	if _, err := db.conn.Exec(sqlRaw, args); err != nil {
		return err
	}
	return nil
}
