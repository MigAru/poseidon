package sqlite

import "github.com/huandu/go-sqlbuilder"

func (db *DB) MarkDeleteRepository(id string) {
	builder := sqlbuilder.NewInsertBuilder()

	builder.InsertInto("repository_delete").Cols("repository_id")
	builder.Values()

}
