package sqlite

import (
	"context"
	"database/sql"
	"github.com/MigAru/poseidon/internal/database/structs"
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
)

func (db *DB) GetRepository(tx *sql.Tx, reference, tag string) (*structs.Repository, error) {
	builder := sqlbuilder.SQLite.NewSelectBuilder()

	builder.Select("id", "reference", "tag", "digest", "created_at", "updated_at")
	builder.From("repository")
	builder.Where(
		builder.Equal("reference", reference),
		builder.Equal("tag", tag),
		builder.Equal("marked", false),
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

func (db *DB) GetRepositoryByID(id string) (*structs.Repository, error) {
	builder := sqlbuilder.SQLite.NewSelectBuilder()

	builder.Select("id", "reference", "tag", "digest", "created_at", "updated_at")
	builder.From("repository")
	builder.Where(builder.Equal("id", id), builder.Equal("marked", false))

	sqlRaw, args := builder.Build()

	row := db.conn.QueryRow(sqlRaw, args...)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var model structs.RepositoryModel
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

func (db *DB) CreateRepository(tx *sql.Tx, project, tag, digest string) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	if tx == nil {
		defaultTx, err := db.conn.BeginTx(context.Background(), nil)

		if err != nil {
			return err
		}
		tx = defaultTx
	}

	builder := sqlbuilder.SQLite.NewInsertBuilder()
	builder.InsertInto("repository").Cols("id", "reference", "tag", "digest")
	builder.Values(id.String(), project, tag, digest)

	sqlRaw, args := builder.Build()

	if _, err := tx.Exec(sqlRaw, args...); err != nil {
		return err
	}

	return nil
}

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

func (db *DB) GetRepositoriesForDelete() ([]*structs.Repository, error) {
	builder := sqlbuilder.SQLite.NewSelectBuilder()
	builder.Select("id", "reference", "tag", "digest", "created_at", "updated_at")
	builder.From("repository")
	builder.Where(builder.Equal("marked", true))

	sqlRaw, args := builder.Build()
	rows, err := db.conn.Query(sqlRaw, args...)
	if err != nil {
		return nil, err
	}

	var res []*structs.Repository
	for rows.Next() {
		var model structs.RepositoryModel
		if err := rows.Scan(
			&model.ID,
			&model.Reference,
			&model.Tag,
			&model.Digest,
			&model.CreatedAt,
			&model.UpdatedAt); err != nil {
			return nil, err
		}

		res = append(res, structs.FromModelToRepository(model))
	}

	return res, nil
}

func (db *DB) DigestUseAnotherRepository() (bool, error) {
	return false, nil
}

func (db *DB) MarkDeleteRepository(tx *sql.Tx, id string) error {
	builder := sqlbuilder.SQLite.NewUpdateBuilder()

	builder.Update("repository").Set(builder.Assign("marked", true))
	builder.Where(builder.Equal("id", id))

	sqlRaw, args := builder.Build()
	if _, err := tx.Exec(sqlRaw, args...); err != nil {
		return err
	}
	return nil
}
