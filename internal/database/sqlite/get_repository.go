package sqlite

import (
	"database/sql"
	"github.com/MigAru/poseidon/internal/database"
)

type RepositoryModel struct {
	reference sql.NullString
	tag       sql.NullString
	createdAt sql.NullTime
	updatedAt sql.NullTime
	attrs     sql.NullString
}
type RepositoryAttrsModel struct {
}

func (db *DB) GetRepository(reference, tag string) (*database.Repository, error) {
	row := db.conn.QueryRow(`
					SELECT 
					    reference, 
					    tag, 
					    created_at, 
					    updated_at, 
					    attrs 
					FROM 
					    repository 
					WHERE 
					    reference = $1 and 
					    tag = $2`, reference, tag)
	if row.Err() != nil {
		return nil, row.Err()
	}
}
