-- +goose Up
CREATE TABLE beef(
    id TEXT PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    meatcut TEXT UNIQUE NOT NULL,
    price FLOAT NOT NULL,
    quantitiy FLOAT
);

-- +goose Down
DROP TABLE beef;