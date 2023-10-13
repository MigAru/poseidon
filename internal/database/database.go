package database

import "time"

type DB interface{}

type Repository struct {
	Reference string
	Tag       string
	CreatedAt time.Time
	UpdatedAt time.Time
	Attrs     RepositoryAttrs
}

type RepositoryAttrs struct {
	Delete bool `json:"delete"`
}
