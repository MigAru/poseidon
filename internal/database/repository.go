package database

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Repository struct {
	ID        string
	Reference string
	Tag       string
	Digest    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Attrs     RepositoryAttrs
}

type RepositoryAttrs struct {
	Delete bool `json:"delete"`
}

func RepositoryAttrsFromRaw(raw string) (RepositoryAttrs, error) {
	var attrs RepositoryAttrs
	err := json.Unmarshal([]byte(raw), &attrs)
	return attrs, err
}

type RepositoryModel struct {
	ID        sql.NullString
	Reference sql.NullString
	Tag       sql.NullString
	Digest    sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Attrs     RepositoryAttrs
}

func FromModelToRepository(model RepositoryModel) *Repository {
	return &Repository{
		ID:        model.ID.String,
		Reference: model.Reference.String,
		Tag:       model.Tag.String,
		Digest:    model.Digest.String,
		CreatedAt: model.CreatedAt.Time,
		UpdatedAt: model.UpdatedAt.Time,
		Attrs:     model.Attrs,
	}
}
