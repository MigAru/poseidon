package sqlite

import (
	"database/sql"
	"github.com/MigAru/poseidon/internal/database"
	"github.com/huandu/go-sqlbuilder"
)

func (db *DB) GetRepository(reference, tag string) (*database.Repository, error) {
	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("reference", "tag", "digest", "created_at", "updated_at", "attrs")
	sb.From("repository")
	sb.Where(sb.Equal("reference", reference), sb.Equal("tag", tag))

	sqlRaw, args := sb.Build()
	row := db.conn.QueryRow(sqlRaw, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var (
		model    database.RepositoryModel
		attrsRaw sql.NullString
	)

	if err := row.Scan(
		&model.Reference,
		&model.Tag,
		&model.Digest,
		&model.CreatedAt,
		&model.UpdatedAt,
		&attrsRaw,
	); err != nil {
		return nil, err
	}

	attrs, err := database.RepositoryAttrsFromRaw(attrsRaw.String)
	if err != nil {
		return nil, err
	}
	model.Attrs = attrs

	return database.FromModelToRepository(model), nil
}
