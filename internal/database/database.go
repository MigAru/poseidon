package database

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MigAru/poseidon/internal/config"
	"github.com/MigAru/poseidon/internal/database/sqlite"
	"github.com/MigAru/poseidon/internal/database/structs"
	_ "github.com/mattn/go-sqlite3"
)

type DB interface {
	NewTx(ctx context.Context) (*sql.Tx, error)
	CreateRepository(tx *sql.Tx, reference, tag, digest string) error
	DeleteRepository(id string) error
	MarkDeleteRepository(id string) error
	UpdateDigestRepository(tx *sql.Tx, project, tag, digest string) error
	GetRepository(tx *sql.Tx, project, tag string) (*structs.Repository, error)
	GetRepositoryByID(id string) (*structs.Repository, error)
}

func New(cfg *config.Config) (DB, func(), error) {
	switch cfg.Database.Driver {
	case SQLite3:
		conn, err := sql.Open(cfg.Database.Driver, cfg.Database.DSN)
		if err != nil {
			return nil, func() {}, err
		}
		return sqlite.New(conn), func() { conn.Close() }, nil
	}

	return nil, func() {}, errors.New("unknown driver")
}
