package sqlite

import (
	"database/sql"
	"github.com/MigAru/poseidon/internal/database"
	"github.com/huandu/go-sqlbuilder"
)

func (db *DB) GetRepository(reference, tag string) (*database.Repository, error) {
	builder := sqlbuilder.NewSelectBuilder()

	builder.Select("r.id", "r.reference", "r.tag", "r.digest", "r.created_at", "r.updated_at")
	builder.From("repository r")
	builder.JoinWithOption(sqlbuilder.RightJoin, "repository_delete rd", "rd.repository_id=r.id")
	builder.Where(builder.Equal("reference", reference), builder.Equal("tag", tag), builder.IsNotNull("rb.id"))

	sqlRaw, args := builder.Build()
	row := db.conn.QueryRow(sqlRaw, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var (
		model    database.RepositoryModel
		attrsRaw sql.NullString
	)

	if err := row.Scan(
		&model.ID,
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
