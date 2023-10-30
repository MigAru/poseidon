package sqlite

import "database/sql"

type DB struct {
	conn *sql.DB
}

func New(conn *sql.DB) *DB {
	return &DB{conn: conn}
}
