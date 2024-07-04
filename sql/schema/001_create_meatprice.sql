-- +goose Up
CREATE TABLE meatprices(
    id TEXT PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    meat_source TEXT NOT NULL,
    meatcut TEXT NOT NULL,
    price FLOAT NOT NULL
);

-- +goose Down
DROP TABLE meatprices;
