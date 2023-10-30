-- +goose Up

CREATE TABLE IF NOT EXISTS repository (
    id varchar(256) NOT NULL UNIQUE,
    reference varchar(512) NOT NULL,
    tag varchar(512) NOT NULL,
    digest varchar(2048) NOT NULL ,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL
);

CREATE OR REPLACE FUNCTION set_current_updated_at()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = current_timestamp;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE OR REPLACE TRIGGER set_current_updated_at
    AFTER UPDATE ON repository
    FOR EACH ROW
    EXECUTE PROCEDURE set_current_updated_at();

CREATE TABLE IF NOT EXISTS repository_delete (
    repository_id varchar(256) NOT NULL UNIQUE
);

-- +goose Down
DROP TRIGGER IF EXISTS set_current_updated_at on repository;
DROP FUNCTION IF EXISTS set_current_updated_at();
DROP TABLE IF EXISTS repository_delete;
DROP TABLE IF EXISTS repository;

