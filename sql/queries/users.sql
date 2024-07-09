-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, username, email, is_admin)
VALUES($1,$2,$3,$4,$5,false)
RETURNING *;

-- name: CreateAdminUser :one
INSERT INTO users(id, created_at, updated_at, username, email, is_admin)
VALUES($1,$2,$3,$4,$5,true)
RETURNING *;

-- name: GetUserID :one
SELECT id FROM users WHERE email = $1;

-- name: StoringLoginToken :exec
UPDATE users SET login_token = $2
WHERE id = $1;

-- name: DeletingLoginToken :exec
UPDATE users SET login_token = NULL
WHERE id = $1;
