CREATE TABLE IF NOT EXISTS auths  (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    token VARCHAR,
    auth_type VARCHAR,
    user_id  BIGINT NOT NULL,
    expires_at TIMESTAMP,
    FOREIGN key (user_id) REFERENCES users(id) ON DELETE CASCADE
);