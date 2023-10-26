package sqlite

import (
	"context"
	"database/sql"
	"github.com/MigAru/poseidon/internal/database/structs"
	"github.com/huandu/go-sqlbuilder"
)

func (db *DB) GetRepository(tx *sql.Tx, project, tag string) (*structs.Repository, error) {
	if tx == nil {
		defaultTx, err := db.conn.BeginTx(context.Background(), nil)
		if err != nil {
			return nil, err
		}
		tx = defaultTx
	}

	builder := sqlbuilder.SQLite.NewSelectBuilder()

	builder.Select("r.id", "r.project", "r.tag", "r.digest", "r.created_at", "r.updated_at")
	builder.From("repository r")
	builder.JoinWithOption(sqlbuilder.LeftJoin, "repository_delete rd", "rd.repository_id=r.id")
	builder.Where(
		builder.Equal("r.project", project),
		builder.Equal("r.tag", tag),
		builder.IsNull("rd.repository_id"),
	)

	sqlRaw, args := builder.Build()
	row := tx.QueryRow(sqlRaw, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var (
		model structs.RepositoryModel
	)

	if err := row.Scan(
		&model.ID,
		&model.Reference,
		&model.Tag,
		&model.Digest,
		&model.CreatedAt,
		&model.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return structs.FromModelToRepository(model), nil
}
