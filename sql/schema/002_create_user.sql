-- +goose Up
CREATE TABLE users(
    id TEXT PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    username TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    is_admin BOOLEAN NOT NULL
);

-- +goose Down
DROP TABLE users;