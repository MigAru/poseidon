package sqlite

import (
	"context"
	"database/sql"
	"github.com/huandu/go-sqlbuilder"
)

func (db *DB) RemoveIndexesLayers(tx *sql.Tx, repositoryID string) error {
	if tx == nil {
		defaultTx, err := db.conn.BeginTx(context.Background(), nil)
		if err != nil {
			return err
		}
		tx = defaultTx
	}

	builder := sqlbuilder.SQLite.NewDeleteBuilder()
	builder.DeleteFrom("repository_layers").Where(builder.Equal("repository_id", repositoryID))

	sqlRaw, args := builder.Build()

	if _, err := tx.Exec(sqlRaw, args...); err != nil {
		return err
	}
	return nil
}

func (db *DB) IndexingLayers(tx *sql.Tx, repositoryID string, layers []int) error {
	if tx == nil {
		defaultTx, err := db.conn.BeginTx(context.Background(), nil)
		if err != nil {
			return err
		}
		tx = defaultTx
	}

	builder := sqlbuilder.SQLite.NewInsertBuilder()
	builder.InsertInto("repository_layers").Cols("repository_id", "digest_id")

	for _, layer := range layers {
		builder.Values(repositoryID, layer)
	}

	sqlRaw, args := builder.Build()
	if _, err := tx.Exec(sqlRaw, args...); err != nil {
		return err
	}

	return nil
}
