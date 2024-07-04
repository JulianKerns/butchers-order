-- +goose Up
CREATE TABLE orders(
    id TEXT PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    user_id TEXT NOT NULL,
    meat_source TEXT NOT NULL,
    meatcut TEXT NOT NULL,
    price FLOAT NOT NULL,
    quantitiy FLOAT NOT NULL,
    UNIQUE(user_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE orders;