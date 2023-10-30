package structs

import (
	"database/sql"
	"time"
)

type Repository struct {
	ID        string
	Reference string
	Tag       string
	Digest    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RepositoryModel struct {
	ID        sql.NullString
	Reference sql.NullString
	Tag       sql.NullString
	Digest    sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

func FromModelToRepository(model RepositoryModel) *Repository {
	return &Repository{
		ID:        model.ID.String,
		Reference: model.Reference.String,
		Tag:       model.Tag.String,
		Digest:    model.Digest.String,
		CreatedAt: model.CreatedAt.Time,
		UpdatedAt: model.UpdatedAt.Time,
	}
}
