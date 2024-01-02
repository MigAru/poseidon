-- +goose Up


CREATE TABLE IF NOT EXISTS repository (
    id varchar(256) NOT NULL UNIQUE,
    reference varchar(512) NOT NULL,
    tag varchar(512) NOT NULL,
    digest varchar(2048) NOT NULL,
    marked BOOLEAN NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT (datetime('now','localtime')),
    updated_at timestamp NOT NULL DEFAULT (datetime('now','localtime'))
);
-- +goose StatementBegin
CREATE TRIGGER IF NOT EXISTS set_current_timestamp_repository_updated_at AFTER UPDATE ON repository
BEGIN
    UPDATE repository SET updated_at=DATETIME('now','localtime') WHERE id = new.id;
END;
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS digests (
    id integer primary key AUTOINCREMENT,
    hash varchar(2048)
);

CREATE TABLE IF NOT EXISTS repository_layers (
    repository_id varchar(256) NOT NULL,
    digest_id integer NOT NULL
);

CREATE TABLE IF NOT EXISTS digest_delete (
  name varchar(2048) NOT NULL
);

-- +goose Down
DROP TRIGGER IF EXISTS set_current_timestamp_repository_updated_at;
DROP TABLE IF EXISTS repository;
DROP TABLE IF EXISTS repositroy_delete;
DROP TABLE IF EXISTS digests;
DROP TABLE IF EXISTS digest_delete;
