CREATE TABLE IF NOT EXISTS currencies (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    code VARCHAR(3) NOT NULL,
    description text,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);