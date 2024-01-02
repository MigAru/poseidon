package sqlite

import (
	"context"
	"database/sql"

	"github.com/huandu/go-sqlbuilder"
)

func (db *DB) IndexingDigest(tx *sql.Tx, digest string) error {
	builder := sqlbuilder.SQLite.NewInsertBuilder()
	builder.InsertIgnoreInto("digests").Cols("hash").Values(digest)

	sqlRaw, args := builder.Build()

	if _, err := tx.Exec(sqlRaw, args...); err != nil {
		return err
	}

	return nil
}

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

func (db *DB) GetDigestsForDelete() ([]string, error) {
	builder := sqlbuilder.SQLite.NewSelectBuilder()
	builder.Select("name").From("digest_delete")

	sqlRaw, args := builder.Build()

	rows, err := db.conn.Query(sqlRaw, args)
	if err != nil {
		return nil, err
	}

	var digests []string
	for rows.Next() {
		var model sql.NullString
		if err := rows.Scan(&model); err != nil {
			return nil, err
		}
		digests = append(digests, model.String)
	}

	return digests, nil
}

func (db *DB) MarkDeleteDigest(digest string) error {
	builder := sqlbuilder.SQLite.NewInsertBuilder()
	builder.InsertIgnoreInto("digest_delete").Cols("name").Values(digest)

	sqlRaw, args := builder.Build()
	if _, err := db.conn.Exec(sqlRaw, args); err != nil {
		return err
	}

	return nil
}
