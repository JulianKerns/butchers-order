-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, username, email, is_admin)
VALUES($1,$2,$3,$4,$5,false)
RETURNING *;

-- name: CreateAdminUser :one
INSERT INTO users(id, created_at, updated_at, username, email, is_admin)
VALUES($1,$2,$3,$4,$5,true)
RETURNING *;
