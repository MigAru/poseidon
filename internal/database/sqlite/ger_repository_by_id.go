package sqlite

import (
	"github.com/MigAru/poseidon/internal/database"
	"github.com/huandu/go-sqlbuilder"
)

func (db *DB) GetRepositoryByID(id string) (*database.Repository, error) {
	builder := sqlbuilder.SQLite.NewSelectBuilder()

	builder.Select("r.id", "r.reference", "r.tag", "r.digest", "r.created_at", "r.updated_at")
	builder.From("repository r")
	builder.JoinWithOption(sqlbuilder.RightJoin, "repository_delete rd", "rd.repository_id=r.id")
	builder.Where(builder.Equal("id", id), builder.IsNotNull("rb.id"))

	sqlRaw, args := builder.Build()

	row := db.conn.QueryRow(sqlRaw, args...)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var model database.RepositoryModel
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

	return database.FromModelToRepository(model), nil
}
